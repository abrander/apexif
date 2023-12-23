package tiff

// IFD is a type representing an Image File Directory in a TIFF file.
type IFD []Entry

func (i IFD) Entry(tag Tag) (Entry, error) {
	for _, entry := range i {
		if Tag(entry.Tag) == tag {
			return entry, nil
		}
	}

	return Entry{}, ErrTagNotFound
}
