package handlers

import (
	"api/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/kuro-helper/core/v2/vndb"
	"github.com/sirupsen/logrus"
)

func SearchBrandHandler(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	result, err := vndb.GetProducerByFuzzy(keyword, "")
	if err != nil {
		logrus.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	// 檢查是否有結果
	if len(result.Producer.Results) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(dto.Response{
			Message: "no results found",
			Data:    nil,
		})
	}

	// 初始化 Link 結構
	linkData := dto.Link{}
	var jawikiUrl string
	for _, link := range result.Producer.Results[0].Extlinks {
		switch link.Name {
		case "website":
			linkData.OfficialWebsite = link.Url
		case "jawiki":
			jawikiUrl = link.Url
		case "enwiki":
			if jawikiUrl == "" {
				linkData.Wikipedia = link.Url
			}
		case "twitter":
			linkData.Xitter = link.Url
		case "steam_curator", "steam":
			linkData.Steam = link.Url
		}
	}
	// 優先使用英文維基
	if jawikiUrl != "" {
		linkData.Wikipedia = jawikiUrl
	}

	// 初始化 VN 陣列
	vnData := make([]dto.VN, 0)
	for _, vn := range result.Vn.Results {
		vnData = append(vnData, dto.VN{
			Title:         vn.Title,
			AltTitle:      vn.Alttitle,
			Average:       vn.Average,
			Rating:        vn.Rating,
			VoteCount:     vn.Votecount,
			LengthMinutes: vn.LengthMinutes,
			LengthVotes:   vn.LengthVotes,
		})
	}

	returnData := dto.SearchBrandResponse{
		ID:          result.Producer.Results[0].ID,
		Name:        result.Producer.Results[0].Name,
		Aliases:     result.Producer.Results[0].Aliases,
		Description: vndb.ConvertBBCodeToMarkdown(result.Producer.Results[0].Description),
		Link:        linkData,
		VN:          vnData,
	}

	return c.Status(fiber.StatusOK).JSON(dto.Response{
		Message: "search successfully",
		Data:    returnData,
	})
}
