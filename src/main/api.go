package main

import (
    "fmt"
    "net/url"
    "net/http"
    "strconv"
    // "strings"
    "encoding/json"
    "io/ioutil"
    "html/template"
    "log"
    // "reflect"

    "github.com/graphql-go/handler"
)

var host string = ""
var port int = 8080
var templates map[string] *template.Template

func main() {
    templates = make(map[string] *template.Template)
    templateLoader("temp", "templates/temp.html")
    templateLoader("start", "templates/start.html")

    InitSessions()
    InitDatabases()
    InitUtils()

    fs := http.FileServer(http.Dir(""))
    http.Handle("/", fs)

    graphqlHandler := handler.New(&handler.Config{Schema: &Schema, Pretty: true})
    http.Handle("/graphql", corsHandler(graphqlHandler))

    http.HandleFunc("/start/", start)
    http.HandleFunc("/endDay", endDay)
    http.HandleFunc("/breedWith", breedWith)
    http.HandleFunc("/setAction", setAction)
    http.HandleFunc("/temp/", tempFunc)
    log.Println("Server up and running on", host + ":" + strconv.Itoa(port))
    log.Println("Go to http://", host + ":" + strconv.Itoa(port) + "/start/ to start a new session")
    http.ListenAndServe(host + ":" + strconv.Itoa(port), nil)
}

func corsHandler(h http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        h.ServeHTTP(w, req)
    }
}

func looking404(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "What are you looking 404?")
}

func start(w http.ResponseWriter, req *http.Request) {
    log.Println("Start")
    context := NewPlaySession("")
    err := templates["start"].Execute(w, context)
    if err != nil { log.Panic(err) }
}

func endDay(w http.ResponseWriter, req *http.Request) {
    // log.Println("End of Day")
    var err error
    asJson, boolParseErr := strconv.ParseBool(getParam(req, "json", 0))

    session := getParam(req, "Session", 0)
    newDay := EndDay(session)
    if boolParseErr == nil && asJson {
        theJson, errJson := json.Marshal(newDay)
        if errJson != nil { log.Panic(errJson) }
        fmt.Fprintf(w, string(theJson))
    } else {
        err = templates["start"].Execute(w, newDay)
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

func setAction(w http.ResponseWriter, req *http.Request) {
    // log.Println("SetAction")
    session := getParam(req, "Session", 0)
    day, _ := strconv.Atoi(getParam(req, "Day", 0))
    creatureID, _ := strconv.Atoi(getParam(req, "CreatureID", 0))
    action := getParam(req, "Action", 0)
    result := SetAction(session, day, creatureID, action)
    fmt.Fprintf(w, result)
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
    obj, foundKey := object[key]
    if foundKey {
        return obj[value]
    }
    return ""
}
