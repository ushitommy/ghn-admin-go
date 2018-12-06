package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/create", createHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//struct for Job JSON data
type Job struct {
	ID    int    `json:"id"`
	Min   string `json:"min"`
	Hour  string `json:"hour"`
	Date  string `json:"date"`
	Month string `json:"month"`
	Days  string `json:"days"`
	Text  string `json:"text"`
}

//トップページ
func rootHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/root.html.tpl"))

	bytes, err := ioutil.ReadFile("data/joblist.json")
	if err != nil {
		log.Fatal(err)
	}

	var job []Job
	if err := json.Unmarshal(bytes, &job); err != nil {
		log.Fatal(err)
	}

	if err := t.ExecuteTemplate(w, "root.html.tpl", job); err != nil {
		log.Fatal(err)
	}
}

//編集ページ
func editHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/edit.html.tpl"))

	if r.Method == "GET" {

		//JSONの読み込み
		bytes, err := ioutil.ReadFile("data/joblist.json")
		if err != nil {
			log.Fatal(err)
		}

		//JSONをパース
		var job []Job
		if err := json.Unmarshal(bytes, &job); err != nil {
			log.Fatal(err)
		}

		//テンプレートを使ってページをレンダリング
		if err := t.ExecuteTemplate(w, "edit.html.tpl", job); err != nil {
			log.Fatal(err)
		}

	} else {

		//JSONの読み込み
		bytes, err := ioutil.ReadFile("data/joblist.json")
		if err != nil {
			log.Fatal(err)
		}

		//バックアップを作成
		if err := ioutil.WriteFile("data/joblist.json.bak", bytes, 0666); err != nil {
			log.Fatal(err)
		}

		//JSONをパース
		var job []Job
		if err := json.Unmarshal(bytes, &job); err != nil {
			log.Fatal(err)
		}

		//Formから受け取ったデータをパース
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		//IDはint型なので変換してから入れる
		fid, _ := strconv.Atoi(r.FormValue("id"))

		//要素の番号として使うため-1する
		editid := fid - 1

		//削除フラグが立っていたら削除処理
		if r.FormValue("delete") == "true" {

			//削除の処理を書く
			//削除というか新しい配列を作って削除フラグが立っていないjobだけ入れるかんじでやるのがやりやすそう @nametakeさん

			//ジョブが1個しかなかったら削除しない（jsonが空だと動かなくなるのでとりあえず）
			if len(job) == 1 {

				//とりあえずページを表示するだけにする
				if err := t.ExecuteTemplate(w, "edit.html.tpl", job); err != nil {
					log.Fatal(err)
				}

			} else {

				//新しい配列を作る
				var ejob []Job

				//削除フラグ付きで送られてきたジョブだけを除外して新しい配列に入れる
				for i := range job {
					if job[i].ID != fid { //Formから送られてきたIDのジョブが削除フラグ付きなのでそれ以外を配列に入れなおす
						newid := job[i].ID
						newmin := job[i].Min
						newhour := job[i].Hour
						newdate := job[i].Date
						newmonth := job[i].Month
						newday := job[i].Days
						newtext := job[i].Text
						newjob := Job{
							ID:    newid,
							Min:   newmin,
							Hour:  newhour,
							Date:  newdate,
							Month: newmonth,
							Days:  newday,
							Text:  newtext,
						}

						//追加を反映
						ejob = append(ejob, newjob)
					}
				}

				//IDが連番にならなくなるので1から振り直す
				for i := range ejob {
					ejob[i].ID = i + 1
				}

				//整形してJSON形式に戻す
				newJSON, err := json.MarshalIndent(ejob, "", "    ")
				if err != nil {
					log.Fatal(err)
				}

				//書き込む
				if err := ioutil.WriteFile("data/joblist.json", newJSON, 0666); err != nil {
					log.Fatal(err)
				}

				//ページをレンダリング
				if err := t.ExecuteTemplate(w, "edit.html.tpl", ejob); err != nil {
					log.Fatal(err)
				}
			}

			//削除フラグなしで来たときは既存ジョブの編集
		} else {

			//Formの値を受け取って要素の値を更新
			job[editid].Min = r.FormValue("min")
			job[editid].Hour = r.FormValue("hour")
			job[editid].Date = r.FormValue("date")
			job[editid].Month = r.FormValue("month")
			job[editid].Days = strings.Join(r.Form["days[]"], ",")
			job[editid].Text = r.FormValue("text")

			//JSON形式にする
			newJSON, err := json.MarshalIndent(job, "", "    ")
			if err != nil {
				log.Fatal(err)
			}

			//JSONファイル上書き
			if err := ioutil.WriteFile("data/joblist.json", newJSON, 0666); err != nil {
				log.Fatal(err)
			}

			//ページ読み込み
			if err := t.ExecuteTemplate(w, "edit.html.tpl", job); err != nil {
				log.Fatal(err)
			}
		}
	}
}

//新規にジョブを追加するページ
func createHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/create.html.tpl"))

	if r.Method == "GET" {

		bytes, err := ioutil.ReadFile("data/joblist.json")
		if err != nil {
			log.Fatal(err)
		}

		var job []Job
		if err := json.Unmarshal(bytes, &job); err != nil {
			log.Fatal(err)
		}

		if err := t.ExecuteTemplate(w, "create.html.tpl", job); err != nil {
			log.Fatal(err)
		}
	} else {

		bytes, err := ioutil.ReadFile("data/joblist.json")
		if err != nil {
			log.Fatal(err)
		}

		//バックアップを作成
		if err := ioutil.WriteFile("data/joblist.json.bak", bytes, 0666); err != nil {
			log.Fatal(err)
		}

		var job []Job
		if err := json.Unmarshal(bytes, &job); err != nil {
			log.Fatal(err)
		}

		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		//フォームから受け取った値を処理して配列に入れる
		newid := len(job) + 1
		newmin := r.FormValue("min")
		newhour := r.FormValue("hour")
		newdate := r.FormValue("date")
		newmonth := r.FormValue("month")
		newday := strings.Join(r.Form["days[]"], ",")
		newtext := r.FormValue("text")
		newjob := Job{
			ID:    newid,
			Min:   newmin,
			Hour:  newhour,
			Date:  newdate,
			Month: newmonth,
			Days:  newday,
			Text:  newtext,
		}

		job = append(job, newjob)

		newJSON, err := json.MarshalIndent(job, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile("data/joblist.json", newJSON, 0666); err != nil {
			log.Fatal(err)
		}

		if err := t.ExecuteTemplate(w, "create.html.tpl", job); err != nil {
			log.Fatal(err)
		}
	}
}
