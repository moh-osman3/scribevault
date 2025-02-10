package diff

type Version struct {
    versionNumber int64
    document *Document
    lastVersion *Version
}

type Document struct {
    name string
    createdTime time.Time
    lastModified time.Time
    body string
    header string
}

type diffFunction func(s1, s2 string) (string, string)

type versionInterface interface {
    getDocumentBody(v *Version) string
    diffVersions(fn diffFunction, v1, v2 *Version) (plusDoc Document, minusDoc Document)
    getLatestVersion(documentName string) *Version
}


