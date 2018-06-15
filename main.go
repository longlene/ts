package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type definitions struct {
	Errno    int    `json:"errno"`
	Errmsg   string `json:"errmsg"`
	BaesInfo struct {
		WordName string `json:"word_name"`
		Exchange struct {
			WordPl    []string      `json:"word_pl"`
			WordPast  []string      `json:"word_past"`
			WordDone  []string      `json:"word_done"`
			WordIng   []string      `json:"word_ing"`
			WordThird []string      `json:"word_third"`
			WordEr    []interface{} `json:"word_er"`
			WordEst   []interface{} `json:"word_est"`
			WordPrep  []interface{} `json:"word_prep"`
			WordAdv   []interface{} `json:"word_adv"`
			WordVerb  []interface{} `json:"word_verb"`
			WordNoun  []interface{} `json:"word_noun"`
			WordAdj   []string      `json:"word_adj"`
			WordConn  []interface{} `json:"word_conn"`
		} `json:"exchange"`
		Symbols []struct {
			PhEn     string `json:"ph_en"`
			PhAm     string `json:"ph_am"`
			PhOther  string `json:"ph_other"`
			PhEnMp3  string `json:"ph_en_mp3"`
			PhAmMp3  string `json:"ph_am_mp3"`
			PhTtsMp3 string `json:"ph_tts_mp3"`
			Parts    []struct {
				Part  string   `json:"part"`
				Means []string `json:"means"`
			} `json:"parts"`
		} `json:"symbols"`
		TranslateType int `json:"translate_type"`
	} `json:"baesInfo"`
	SameAnalysis []struct {
		PartName string   `json:"part_name"`
		WordList string   `json:"word_list"`
		Means    []string `json:"means"`
	} `json:"sameAnalysis"`
	TradeMeans []struct {
		WordTrade string   `json:"word_trade"`
		WordMean  []string `json:"word_mean"`
	} `json:"trade_means"`
	Sentence []struct {
		NetworkID   string `json:"Network_id"`
		NetworkEn   string `json:"Network_en"`
		NetworkCn   string `json:"Network_cn"`
		TtsMp3      string `json:"tts_mp3"`
		TtsSize     string `json:"tts_size"`
		SourceType  int    `json:"source_type"`
		SourceID    int    `json:"source_id"`
		SourceTitle string `json:"source_title"`
	} `json:"sentence"`
	Netmean struct {
		PerfectNetExp []struct {
			ID  string `json:"id"`
			Key string `json:"key"`
			Exp string `json:"exp"`
			URL string `json:"url"`
			Bas int    `json:"bas"`
			Abs string `json:"abs"`
		} `json:"PerfectNetExp"`
		RelatedPhrase []struct {
			Word string `json:"word"`
			List []struct {
				ID  string `json:"id"`
				Key string `json:"key"`
				Exp string `json:"exp"`
				URL string `json:"url"`
				Bas int    `json:"bas"`
				Abs string `json:"abs"`
			} `json:"list"`
		} `json:"RelatedPhrase"`
	} `json:"netmean"`
	Phrase []struct {
		CizuName string `json:"cizu_name"`
		Jx       []struct {
			JxEnMean string        `json:"jx_en_mean"`
			JxCnMean string        `json:"jx_cn_mean"`
			Lj       []interface{} `json:"lj"`
		} `json:"jx"`
	} `json:"phrase"`
	Synonym []struct {
		PartName string `json:"part_name"`
		Means    []struct {
			WordMean string   `json:"word_mean"`
			Cis      []string `json:"cis"`
		} `json:"means"`
	} `json:"synonym"`
	Encyclopedia struct {
		Image   string `json:"image"`
		URL     string `json:"url"`
		Content string `json:"content"`
	} `json:"encyclopedia"`
	Collins []struct {
		Entry []struct {
			Posp    string `json:"posp"`
			Tran    string `json:"tran"`
			Def     string `json:"def"`
			Example []struct {
				Ex      string `json:"ex"`
				Tran    string `json:"tran"`
				TtsMp3  string `json:"tts_mp3"`
				TtsSize string `json:"tts_size"`
			} `json:"example"`
		} `json:"entry"`
	} `json:"collins"`
	EeMean []struct {
		PartName string `json:"part_name"`
		Means    []struct {
			WordMean  string `json:"word_mean"`
			Sentences []struct {
				Sentence string `json:"sentence"`
			} `json:"sentences"`
		} `json:"means"`
	} `json:"ee_mean"`
	AuthSentence []struct {
		ID            string `json:"id"`
		Content       string `json:"content"`
		Link          string `json:"link"`
		ShortLink     string `json:"short_link"`
		Source        string `json:"source"`
		Score         string `json:"score"`
		CacheStatus   string `json:"cache_status"`
		TtsMp3        string `json:"tts_mp3"`
		TtsSize       string `json:"tts_size"`
		Diff          string `json:"diff"`
		Oral          string `json:"oral"`
		ResContent    string `json:"res_content"`
		ResContentCon string `json:"res_content_con"`
		ResKey        string `json:"res_key"`
		SourceType    int    `json:"source_type"`
		SourceID      int    `json:"source_id"`
		SourceTitle   string `json:"source_title"`
	} `json:"auth_sentence"`
	CetFour struct {
		Word     string        `json:"word"`
		Count    int           `json:"count"`
		Kd       []interface{} `json:"kd"`
		Sentence []struct {
			Sentence string `json:"sentence"`
			Come     string `json:"come"`
		} `json:"Sentence"`
	} `json:"cetFour"`
	Bidec struct {
		WordName string `json:"word_name"`
		Parts    []struct {
			PartName string `json:"part_name"`
			WordID   string `json:"word_id"`
			PartID   string `json:"part_id"`
			Means    []struct {
				WordMean  string `json:"word_mean"`
				PartID    string `json:"part_id"`
				MeanID    string `json:"mean_id"`
				Sentences []struct {
					En string `json:"en"`
					Cn string `json:"cn"`
				} `json:"sentences"`
			} `json:"means"`
		} `json:"parts"`
	} `json:"bidec"`
	Jushi []struct {
		English string `json:"english"`
		Chinese string `json:"chinese"`
		Mp3     string `json:"mp3"`
	} `json:"jushi"`
	WordFlag  int      `json:"_word_flag"`
	Exchanges []string `json:"exchanges"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ts word")
		os.Exit(1)
	}

	data, err := query(os.Args[1])
	if err != nil {
		panic(err)
	}
	display(data)
}

func query(word string) ([]byte, error) {
	const URL string = "http://www.iciba.com/index.php"

	parameter := url.Values{}
	parameter.Set("a", "getWordMean")
	parameter.Set("c", "search")
	parameter.Set("list", "1,4")
	parameter.Set("word", word)

	var ur string = URL + "?" + parameter.Encode()
	resp, err := http.Get(ur)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func display(data []byte) {
	var definition = new(definitions)
	//fmt.Println("data:%s\n", string(data))

	err := json.Unmarshal(data, &definition)
	if err != nil {
		fmt.Println("[E]: ", err)
	}

	if definition.Errno != 0 {
		fmt.Println("[E] errno: ", definition.Errno)
	}

	fmt.Printf("[\033[1;32m%s\033[0m]\n", definition.BaesInfo.WordName)
	for _, eemean := range definition.EeMean {
		fmt.Printf("\033[3;33m%s\033[0m\n", eemean.PartName)
		for _, mean := range eemean.Means {
			fmt.Printf("\033[36m%s\033[0m\n", mean.WordMean)
		}
	}
}
