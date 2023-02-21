package core

import (
	"fmt"
	"path/filepath"
	"sardines/config"
	se "sardines/core/searchable_encryption"
	"sardines/err"
	"sardines/storage"
	"sardines/tool"
	"strings"
)

func (h *HostNode) UploadFile(filePath, name, index string) (string, error) {

	// encrypt the content
	cp, ip, e := se.Upload(filePath, name, index)
	if e != nil {
		return "", e
	}
	content, e := tool.LoadFile(cp)
	if e != nil {
		return "", e
	}
	file := tool.NewFileFromContent(fmt.Sprintf("%s:%s", name, index), content)

	// upload file to ipfs
	cid, err2 := h.api.Upload(file.Raw())
	if err2 != nil {
		return "", err2
	}
	file.CID = cid

	// upload the inverted index
	content, e = tool.LoadFile(ip)
	if e != nil {
		return "", e
	}
	indexID, e := h.api.Upload(content)
	if e != nil {
		return "", e
	}

	fmt.Println("上传的倒排索引文件为：", indexID)

	// update the keyTable
	if b := h.Ktab.Append(indexID, []string{file.CID}); !b {
		return "", err.KeyTableUpdateErr
	}

	// update the manifest
	err2 = storage.UpdateManifest(file.Entry)
	if err2 != nil {
		return "", err2
	}

	return file.ID(), nil
}

func (h *HostNode) SearchFileByCid(cid string) (*tool.File, error) {
	bytes, e := h.api.Download(cid)
	if e != nil {
		return nil, e
	}

	file, e := tool.NewFileFromRaw(bytes)
	if e != nil {
		return nil, e
	}
	file.CID = cid

	e = storage.UpdateManifest(file.Entry)
	if e != nil {
		return file, e
	}
	return file, nil
}

func (h *HostNode) SearchFileByKey(kw string) ([]*tool.File, error) {
	res := make([]*tool.File, 0, 5)

	// 生成用户查询Token
	err := se.GenToken(kw)
	if err != nil {
		return nil, err
	}

	indexPath := filepath.Join(config.WD, "SerializedData/SearchableEncryption/InvertedIndex.ser")

	keys := h.Ktab.GetAllKeys()

	var files []string
	// 倒排索引匹配
	for _, key := range keys {
		bytes, err2 := h.api.Download(key)
		if err2 != nil {
			continue
		}
		err2 = tool.WriteFile(bytes, indexPath)
		if err2 != nil {

			continue
		}

		if b := se.NodeSearch(); b {
			file := h.Ktab.Get(key)
			files = append(files, file...)
			fmt.Println("匹配的倒排索引文件为：", key)
			//fmt.Println(key, files)
			//break
		}
	}

	// 下载加密文件, 并解密
	for _, id := range files {
		file, err1 := h.SearchFileByCid(id)
		if err1 != nil {
			continue
		}
		storage.DownloadFile(file)
		name := strings.Split(file.Origin, ":")[0]
		se.DecFile(name)

		bytes, _ := tool.LoadFile(filepath.Join(config.WD, "PlainFiles", name))
		file.Content = bytes
		res = append(res, file)
	}
	return res, nil
}
