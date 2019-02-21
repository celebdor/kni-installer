package bootkube

import (
	"os"
	"path/filepath"

	"github.com/metalkube/kni-installer/pkg/asset"
	"github.com/metalkube/kni-installer/pkg/asset/templates/content"
)

const (
	kubeSystemConfigmapEtcdCAFileName = "kube-system-configmap-etcd-ca-bundle.yaml.template"
)

var _ asset.WritableAsset = (*KubeSystemConfigmapEtcdCA)(nil)

// KubeSystemConfigmapEtcdCA is the constant to represent contents of kube-system-configmap-etcd-ca-bundle.yaml.template file.
type KubeSystemConfigmapEtcdCA struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *KubeSystemConfigmapEtcdCA) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *KubeSystemConfigmapEtcdCA) Name() string {
	return "KubeSystemConfigmapEtcdCA"
}

// Generate generates the actual files by this asset
func (t *KubeSystemConfigmapEtcdCA) Generate(parents asset.Parents) error {
	fileName := kubeSystemConfigmapEtcdCAFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *KubeSystemConfigmapEtcdCA) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *KubeSystemConfigmapEtcdCA) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, kubeSystemConfigmapEtcdCAFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
