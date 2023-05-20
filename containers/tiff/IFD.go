package tiff

type IFD []Entry

func (i IFD) Entry(tag Tag) (Entry, error) {
	for _, entry := range i {
		if Tag(entry.Tag) == tag {
			return entry, nil
		}
	}

	return Entry{}, ErrTagNotFound
}
