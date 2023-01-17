package handler

import "incrediblefour/features/content"

type contentHandle struct {
	srv content.ContentService
}

func New(cs content.ContentService) content.ContentHandler {
	return &contentHandle{
		srv: cs,
	}
}
