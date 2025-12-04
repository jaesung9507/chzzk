package chzzk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LiveStatusResp struct {
	Code    int `json:"code"`
	Message any `json:"message"`
	Content struct {
		LiveTitle                     string   `json:"liveTitle"`
		Status                        string   `json:"status"`
		ConcurrentUserCount           int      `json:"concurrentUserCount"`
		CvExposure                    bool     `json:"cvExposure"`
		AccumulateCount               int      `json:"accumulateCount"`
		PaidPromotion                 bool     `json:"paidPromotion"`
		Adult                         bool     `json:"adult"`
		KrOnlyViewing                 bool     `json:"krOnlyViewing"`
		OpenDate                      string   `json:"openDate"`
		CloseDate                     *string  `json:"closeDate"`
		ClipActive                    bool     `json:"clipActive"`
		ChatChannelId                 string   `json:"chatChannelId"`
		Tags                          []string `json:"tags"`
		CategoryType                  string   `json:"categoryType"`
		LiveCategory                  string   `json:"liveCategory"`
		LiveCategoryValue             string   `json:"liveCategoryValue"`
		LivePollingStatusJson         string   `json:"livePollingStatusJson"`
		FaultStatus                   any      `json:"faultStatus"`
		UserAdultStatus               *string  `json:"userAdultStatus"`
		AbroadCountry                 bool     `json:"abroadCountry"`
		BlindType                     any      `json:"blindType"`
		PlayerRecommendContent        any      `json:"playerRecommendContent"`
		ChatActive                    bool     `json:"chatActive"`
		ChatAvailableGroup            string   `json:"chatAvailableGroup"`
		ChatAvailableCondition        string   `json:"chatAvailableCondition"`
		MinFollowerMinute             int      `json:"minFollowerMinute"`
		AllowSubscriberInFollowerMode bool     `json:"allowSubscriberInFollowerMode"`
		ChatSlowModeSec               int      `json:"chatSlowModeSec"`
		ChatEmojiMode                 bool     `json:"chatEmojiMode"`
		ChatDonationRankingExposure   bool     `json:"chatDonationRankingExposure"`
		DropsCampaignNo               any      `json:"dropsCampaignNo"`
		LiveTokenList                 []string `json:"liveTokenList"`
		WatchPartyNo                  int      `json:"watchPartyNo"`
		WatchPartyTag                 string   `json:"watchPartyTag"`
		TimeMachineActive             bool     `json:"timeMachineActive"`
		ChannelId                     string   `json:"channelId"`
		LastAdultReleaseDate          any      `json:"lastAdultReleaseDate"`
		LastKrOnlyViewingReleaseDate  any      `json:"lastKrOnlyViewingReleaseDate"`
		LastTvAppAllowedDate          any      `json:"lastTvAppAllowedDate"`
		TvAppViewingPolicyType        string   `json:"tvAppViewingPolicyType"`
		LogPowerActive                bool     `json:"logPowerActive"`
		LogPowerRankingExposure       bool     `json:"logPowerRankingExposure"`
	} `json:"content"`
}

func (c *Client) GetLiveStatus(channelID string) (*LiveStatusResp, error) {
	url := fmt.Sprintf("https://api.chzzk.naver.com/polling/v3.1/channels/%s/live-status", channelID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}
	c.header["Referer"] = fmt.Sprintf("https://chzzk.naver.com/live/%s", channelID)

	resp, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to reuqest: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	result := &LiveStatusResp{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return result, nil
}

type LivePlaybackInfo struct {
	Meta struct {
		VideoId   string `json:"videoId"`
		StreamSeq int64  `json:"streamSeq"`
		LiveId    string `json:"liveId"`
		PaidLive  bool   `json:"paidLive"`
		CDNInfo   struct {
			CDNType string `json:"cdnType"`
		} `json:"cdnInfo"`
		CmcdEnabled      bool   `json:"cmcdEnabled"`
		PlaybackAuthType string `json:"playbackAuthType"`
	} `json:"meta"`
	ServiceMeta struct {
		ContentType string `json:"contentType"`
	} `json:"serviceMeta"`
	Live struct {
		Start       string `json:"start"`
		Open        string `json:"open"`
		TimeMachine bool   `json:"timeMachine"`
		Status      string `json:"status"`
	} `json:"live"`
	API []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"api"`
	Media []struct {
		MediaId       string `json:"mediaId"`
		Protocol      string `json:"protocol"`
		Path          string `json:"path"`
		Latency       string `json:"latency,omitempty"`
		EncodingTrack []struct {
			EncodingTrackId    string `json:"encodingTrackId"`
			Path               string `json:"path,omitempty"`
			VideoProfile       string `json:"videoProfile,omitempty"`
			AudioProfile       string `json:"audioProfile,omitempty"`
			P2PPath            string `json:"p2pPath,omitempty"`
			P2PPathUrlEncoding string `json:"p2pPathUrlEncoding,omitempty"`
			VideoCodec         string `json:"videoCodec,omitempty"`
			VideoBitRate       int64  `json:"videoBitRate,omitempty"`
			AudioCodec         string `json:"audioCodec,omitempty"`
			AudioBitRate       int64  `json:"audioBitRate"`
			AudioOnly          bool   `json:"audioOnly,omitempty"`
			VideoFrameRate     string `json:"videoFrameRate,omitempty"`
			VideoWidth         int    `json:"videoWidth,omitempty"`
			VideoHeight        int    `json:"videoHeight,omitempty"`
			AudioSamplingRate  int64  `json:"audioSamplingRate"`
			AudioChannel       int    `json:"audioChannel"`
			AvoidReencoding    bool   `json:"avoidReencoding"`
			VideoDynamicRange  string `json:"videoDynamicRange,omitempty"`
		} `json:"encodingTrack"`
	} `json:"media"`
	Thumbnail struct {
		SnapshotThumbnailTemplate string `json:"snapshotThumbnailTemplate"`
		SpriteSeekingThumbnail    struct {
			SpriteFormat struct {
				RowCount        int    `json:"rowCount"`
				ColumnCount     int    `json:"columnCount"`
				IntervalType    string `json:"intervalType"`
				Interval        int    `json:"interval"`
				ThumbnailWidth  int    `json:"thumbnailWidth"`
				ThumbnailHeight int    `json:"thumbnailHeight"`
			} `json:"spriteFormat"`
			URLTemplate       string `json:"urlTemplate"`
			ProcessingSeconds int    `json:"processingSeconds"`
		} `json:"spriteSeekingThumbnail"`
		Types []string `json:"types"`
	} `json:"thumbnail"`
	Multiview []struct{} `json:"multiview"`
}

type LiveDetailResp struct {
	Code    int `json:"code"`
	Message any `json:"message"`
	Content struct {
		LiveID                        int64    `json:"liveId"`
		LiveTitle                     string   `json:"liveTitle"`
		Status                        string   `json:"status"`
		LiveImageUrl                  string   `json:"liveImageUrl"`
		DefaultThumbnailImageUrl      any      `json:"defaultThumbnailImageUrl"`
		ConcurrentUserCount           int      `json:"concurrentUserCount"`
		CvExposure                    bool     `json:"cvExposure"`
		AccumulateCount               int      `json:"accumulateCount"`
		OpenDate                      string   `json:"openDate"`
		CloseDate                     *string  `json:"closeDate"`
		Adult                         bool     `json:"adult"`
		KrOnlyViewing                 bool     `json:"krOnlyViewing"`
		ClipActive                    bool     `json:"clipActive"`
		Tags                          []string `json:"tags"`
		ChatChannelId                 string   `json:"chatChannelId"`
		CategoryType                  string   `json:"categoryType"`
		LiveCategory                  string   `json:"liveCategory"`
		LiveCategoryValue             string   `json:"liveCategoryValue"`
		ChatActive                    bool     `json:"chatActive"`
		ChatAvailableGroup            string   `json:"chatAvailableGroup"`
		PaidPromotion                 bool     `json:"paidPromotion"`
		ChatAvailableCondition        string   `json:"chatAvailableCondition"`
		MinFollowerMinute             int      `json:"minFollowerMinute"`
		AllowSubscriberInFollowerMode bool     `json:"allowSubscriberInFollowerMode"`
		LivePlaybackJson              string   `json:"livePlaybackJson"`
		P2PQuality                    []string `json:"p2pQuality"`
		TimeMachineActive             bool     `json:"timeMachineActive"`
		TimeMachinePlayback           bool     `json:"timeMachinePlayback"`
		Channel                       any      `json:"channel"`
		LivePollingStatusJson         string   `json:"livePollingStatusJson"`
		UserAdultStatus               *string  `json:"userAdultStatus"`
		BlindType                     any      `json:"blindType"`
		ChatDonationRankingExposure   bool     `json:"chatDonationRankingExposure"`
		LogPower                      any      `json:"logPower"`
		AdParameter                   any      `json:"adParameter"`
		DropsCampaignNo               any      `json:"dropsCampaignNo"`
		WatchPartyNo                  int      `json:"watchPartyNo"`
		WatchPartyTag                 string   `json:"watchPartyTag"`
		WatchPartyPaidProductId       any      `json:"watchPartyPaidProductId"`
		Dab                           bool     `json:"dab"`
		Earthquake                    bool     `json:"earthquake"`
		PaidProduct                   any      `json:"paidProduct"`
		TvAppViewingPolicyType        string   `json:"tvAppViewingPolicyType"`
		Party                         any      `json:"party"`
	} `json:"content"`
}

func (i *LivePlaybackInfo) GetHLSPath() string {
	for _, media := range i.Media {
		if media.MediaId == "HLS" {
			return media.Path
		}
	}

	return ""
}

func (r *LiveDetailResp) GetLivePlayback() (*LivePlaybackInfo, error) {
	result := &LivePlaybackInfo{}
	if err := json.Unmarshal([]byte(r.Content.LivePlaybackJson), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetLiveDetail(channelID string) (*LiveDetailResp, error) {
	url := fmt.Sprintf("https://api.chzzk.naver.com/service/v3.2/channels/%s/live-detail", channelID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}
	c.header["Referer"] = fmt.Sprintf("https://chzzk.naver.com/live/%s", channelID)

	resp, err := c.do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to reuqest: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	result := &LiveDetailResp{}
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return result, nil
}
