package handler

import (
	"context"
	"strconv"

	"github.com/duchoang206h/telebot-storage/bot"
	"github.com/duchoang206h/telebot-storage/config"
	"github.com/duchoang206h/telebot-storage/proto"
)
type GRPCHandler struct {
	proto.UnimplementedFileStorageServer
}

func (h *GRPCHandler) UploadFile(ctx context.Context, req *proto.UploadFileRequest) (*proto.FileResponse, error) {
	chat_id, _ := strconv.Atoi(config.Config("CHAT_ID"))
	telebot := bot.GetBot()
	fileBytes, fileName:= req.Content, req.Name
	fileId, err := telebot.UploadFile(fileBytes, fileName, int64(chat_id))
	if err != nil {
		return nil, err
	}
	return &proto.FileResponse{ Result: &proto.FileResult{Result: fileId}, }, nil
}
func (h *GRPCHandler) GetFile(ctx context.Context, req *proto.GetFileRequest) (*proto.FileResponse, error) {
	telebot := bot.GetBot()
	fileUrl, err := telebot.GetFileUrl(req.FileId)
	if err != nil {
		return nil, err
	}
	return &proto.FileResponse{ Result: &proto.FileResult{ Result: fileUrl }}, nil
}