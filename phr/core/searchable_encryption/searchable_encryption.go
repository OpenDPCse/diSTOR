package searchable_encryption

import (
	"os/exec"
	"path/filepath"
	"sardines/config"
	"strings"
)

var (
	uploadJar     string
	tokenJar      string
	NodeSearchJar string
	DecFilesJar   string
)

func init() {
	uploadJar = filepath.Join(config.WD, "javac\\FilesUploader.jar")
	tokenJar = filepath.Join(config.WD, "javac\\FilesSearcher.jar")
	NodeSearchJar = filepath.Join(config.WD, "javac\\NodeSearch.jar")
	DecFilesJar = filepath.Join(config.WD, "javac\\DecFiles.jar")
}

func Upload(path, name, key string) (cypherPath string, indexPath string, er error) {

	cypherPath = filepath.Join(config.WD, "UploadFiles", name)
	indexPath = filepath.Join(config.WD, "SerializedData/SearchableEncryption/InvertedIndex.ser")

	cmd := exec.Command("java", "-jar", uploadJar, "-a", "ECNU teacher doctor master bachelor 2016 2015 2014", "-p", "ECNU teacher", "-i", key, "-f", path)

	_, err := cmd.Output()
	// fmt.Println(string(out))
	if err != nil {
		return "", "", err
	}
	return cypherPath, indexPath, nil
}

func GenToken(kw string) error {

	cmd := exec.Command("java", "-jar", tokenJar, "-a", "ECNU teacher doctor master bachelor 2016 2015 2014", "-s", "ECNU teacher", "-w", kw)
	_, err := cmd.Output()
	//fmt.Println(string(out))
	if err != nil {
		return err
	}

	return nil
}

func NodeSearch() bool {
	cmd := exec.Command("java", "-jar", NodeSearchJar)
	output, err := cmd.Output()
	if err != nil || strings.Index(string(output), "not") >= 0 {
		return false
	} else if strings.Index(string(output), "have") >= 0 {
		return true
	} else {
		return false
	}
}

func DecFile(name string) error {
	cmd := exec.Command("java", "-jar", DecFilesJar, "-s", "ECNU teacher", "-e", filepath.Join(config.Downloads, name))
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
