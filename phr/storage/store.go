package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sardines/config"
	"sardines/tool"
	"strings"
)

var manifest map[string]string

func init() {
	manifest = make(map[string]string)
	exist, er := tool.HasDir(config.Manifest)
	if exist {
		file, e := tool.LoadFile(config.Manifest)
		if e != nil {
			return
		}
		if len(file) == 0 {
			return
		}
		e = json.Unmarshal(file, &manifest)
		if e != nil {
			return
		}
	} else if !exist && er == nil {
		os.Create(config.Manifest)
		return
	}
}

func FileStoreTree() map[string][]string {
	res := make(map[string][]string)
	res[""] = []string{}
	for cid, origin := range manifest {
		res[""] = append(res[""], cid)
		res[cid] = append(res[cid], origin)

	}

	return res
}

func UpdateManifest(entry tool.Entry) error {
	exist, err2 := tool.HasDir(config.Manifest)
	if err2 != nil {
		return err2
	}
	if exist {
		file, err2 := tool.LoadFile(config.Manifest)
		if err2 != nil {
			return err2
		}
		if len(file) != 0 {
			mf := make(map[string]string, 0)
			err2 = json.Unmarshal(file, &mf)
			if err2 != nil {
				return err2
			}
			mf[entry.CID] = entry.Origin
			j, err2 := json.Marshal(mf)
			if err2 != nil {
				return err2
			}
			er := tool.WriteFile(j, config.Manifest)
			if er != nil {
				return er
			}
			manifest = mf
			return nil
		} else {
			mf := make(map[string]string)
			mf[entry.CID] = entry.Origin
			j, err2 := json.Marshal(mf)
			if err2 != nil {
				return err2
			}
			er := tool.WriteFile(j, config.Manifest)
			if er != nil {
				return er
			}
			manifest = mf
			return nil
		}
	} else {
		mf := make(map[string]string)
		mf[entry.CID] = entry.Origin
		j, err2 := json.Marshal(mf)
		if err2 != nil {
			return err2
		}
		er := tool.WriteFile(j, config.Manifest)
		if er != nil {
			return er
		}
		manifest = mf
		return nil
	}
}

func OriginalNameFromFID(fid string) string {
	return manifest[fid]
}

// DownloadFile 下载文件
func DownloadFile(file *tool.File) error {
	name := strings.Split(file.Origin, ":")[0]
	filePath := filepath.Join(config.Downloads, name)
	er := tool.WriteFile(file.Content, filePath)
	if er != nil {
		return er
	}
	return nil
}
