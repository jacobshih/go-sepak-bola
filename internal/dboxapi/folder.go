package dboxapi

/*
 * FileEntry
 */
type FileEntry struct {
	Name     string
	Path     string
	Modified string
}

func (file *FileEntry) GetName() string {
	return file.Name
}

func (file *FileEntry) GetPath() string {
	return file.Path
}

func (file *FileEntry) GetModified() string {
	return file.Modified
}

/*
 * FileEntries
 */
type FileEntries []FileEntry

func (files *FileEntries) Add(name string, path string, modified string) FileEntries {
	if files == nil {
		*files = FileEntries{}
	}
	entry := FileEntry{Name: name, Path: path, Modified: modified}
	*files = append(*files, entry)
	return *files
}

func (files *FileEntries) ClearFiles() FileEntries {
	*files = nil
	return *files
}

/*
 * Folder
 */
type Folder struct {
	Files FileEntries
	Link  string
}

func (folder *Folder) GetFiles() *FileEntries {
	return &folder.Files
}

func (folder *Folder) GetLink() string {
	return folder.Link
}

/*
 * Folders
 */
type Folders map[string]Folder

/*
 * FIXME: should it be renamed to Get() for more suitable?
 */
func (folders *Folders) Add(path string) (folder Folder) {
	if *folders == nil {
		*folders = Folders{}
	}
	if _, found := (*folders)[path]; !found {
		(*folders)[path] = Folder{}
	}
	return (*folders)[path]
}

func (folders *Folders) Clear() {
	for k := range *folders {
		delete(*folders, k)
	}
}
