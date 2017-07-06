package main

import (
    "fmt"
    "net/url"
    "net/http"
    "strconv"
    // "strings"
    // "encoding/json"
    "io/ioutil"
    "html/template"
    "log"
    // "reflect"
)

var host string = ""
var port int = 8080
var templates map[string] *template.Template

func main() {
    templates = make(map[string] *template.Template)
    templateLoader("temp", "templates/temp.html")
    templateLoader("start", "templates/start.html")

    InitSessions()
    fs := http.FileServer(http.Dir(""))
    http.Handle("/", fs)
    http.HandleFunc("/start/", start)
    http.HandleFunc("/endDay", endDay)
    http.HandleFunc("/breedWith", breedWith)
    http.HandleFunc("/temp/", tempFunc)
    log.Println("Server up and running on", host + ":" + strconv.Itoa(port))
    http.ListenAndServe(host + ":" + strconv.Itoa(port), nil)
}

func looking404(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "What are you looking 404?")
}

func start(w http.ResponseWriter, req *http.Request) {
    log.Println("Start")
    err := templates["start"].Execute(w, NewPlaySession(""))
    if err != nil { log.Panic(err) }
}

func endDay(w http.ResponseWriter, req *http.Request) {
    // log.Println("End of Day")
    session := getParam(req, "Session", 0)
    newDay := EndDay(session)
    if newDay.Play.Resources < 0 {
        err := templates["start"].Execute(w, NewPlaySession(session))
        if err != nil { log.Panic(err) }
    } else {
        err := templates["start"].Execute(w, newDay)
        if err != nil { log.Panic(err) }
    }
}

func breedWith(w http.ResponseWriter, req *http.Request) {
    log.Println("BreedWith")
    session := getParam(req, "Session", 0)
    creature1ID, _ := strconv.Atoi(getParam(req, "Creature1ID", 0))
    creature2ID, _ := strconv.Atoi(getParam(req, "Creature2ID", 0))
    result := BreedWith(session, creature1ID, creature2ID)
    fmt.Fprintf(w, strconv.FormatBool(result))
}

func tempFunc(w http.ResponseWriter, req *http.Request) {
    templates["temp"].Execute(w, nil)
}

func templateLoader(name string, templateFile string) {
    stream, err := ioutil.ReadFile(templateFile)
    if err != nil { log.Panic(err) }
    // fmt.Println(string(stream))
    tempTemplate, err := template.New(name).Parse(string(stream))
    if err != nil { log.Panic(err) }
    templates[name] = tempTemplate
}

func getParam(req *http.Request, key string, value int) string {
    req.ParseForm()
    object, err := url.ParseQuery(req.URL.RawQuery)
    if err != nil { log.Panic(err) }
    return object[key][value]
}


// from flask import Flask
// from flask import render_template
// import requests
// import ast
//
// app = Flask(__name__)
//
// ##############################
// # API contact                #
// ##############################
//
// creatures_api_url = "http://127.0.0.1:5000"
//
//
// def new_session_play():
//     return requests.get(creatures_api_url + "/newPlaySession/").content
//
//
// def get_session_play(session):
//     contents = requests.get(creatures_api_url + "/getPlayState/" + session).content
//     play_dict = ast.literal_eval(contents)
//     creatures = [creature for creature in play_dict['creatures'].values()]
//     creatures = sorted(creatures, key=lambda creature_sort: creature_sort['ident'])
//     play_dict['creatures'] = creatures
//     return play_dict
//
//
// def end_day_for_session_play(session):
//     contents = requests.post(creatures_api_url + "/endDay/" + session + "/").content
//     return ast.literal_eval(contents)
//
//
// ##############################
// # Entry points               #
// ##############################
//
//
// @app.route('/start/', methods=['GET'])
// def start():
//     print("Entered Start")
//     session = new_session_play()
//     return instructions_page(session, get_session_play(session))
//
//
// @app.route('/instructions/<session>/', methods=['GET'])
// def instructions(session):
//     print("Entered Instructions")
//     return instructions_page(session, get_session_play(session))
//
//
// @app.route('/breeding/<session>/', methods=['GET'])
// def breeding(session):
//     print("Entered Breeding")
//     return breeding_page(session, get_session_play(session))
//
//
// @app.route('/training/<session>/', methods=['GET'])
// def training(session):
//     print("Entered Working")
//     return training_page(session, get_session_play(session))
//
//
// @app.route('/endDay/<session>/', methods=['GET'])
// def end_day(session):
//     print("Entered EndDay")
//     state = end_day_for_session_play(session)
//     if not state['success']:
//         message = "Game Restarted: No more resources"
//         return instructions_page(session, get_session_play(session), message)
//     return instructions_page(session, get_session_play(session))
//
//
// ##############################
// # pages to render            #
// ##############################
//
//
// def instructions_page(session, current_play, message=""):
//     context = get_context(session, message, "Instruction Options", current_play)
//     return render_template('instructions.html', context=context)
//
//
// def breeding_page(session, current_play, message=""):
//     context = get_context(session, message, "Breeding Options", current_play)
//     return render_template('breeding.html', context=context)
//
//
// def training_page(session, current_play, message=""):
//     context = get_context(session, message, "Training Options", current_play)
//     return render_template('training.html', context=context)
//
//
// def get_context(session, message, header, current_play):
//     return {'session': session, 'message': message, 'header': header, 'play': current_play}
//
//
// if __name__ == '__main__':
//     app.debug = True
//     app.run(host='0.0.0.0', port=5001)
