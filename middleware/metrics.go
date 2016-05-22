package middleware

import (
    "github.com/prometheus/client_golang/prometheus"
    dto "github.com/prometheus/client_model/go"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/mutil"
    "github.com/golang/protobuf/proto"
    "net/http"
    "time"
    "os"
    "io"
    "fmt"
)

var h = prometheus.NewHistogram(
    prometheus.HistogramOpts{
        Name:        "response_times",
        Help:        "auto",
        //ConstLabels: s.labels,
        Buckets: []float64{.5, 1, 5, 10, 12, 15, 17.5, 20, 22, 25, 30, 40, 50, 100, 500, 1000},
        //Buckets: prometheus.LinearBuckets(20, 5, 5),
        //Buckets:     prometheus.LinearBuckets(start, width, count),
    },
)

type Monitor struct {
    //histograms   map[string]prometheus.Histogram
    h prometheus.Histogram
}

var M = Monitor{h:h}

func (m *Monitor) Write(w http.ResponseWriter) {
    metric := &dto.Metric{}
    m.h.Write(metric)
    b := []byte(proto.MarshalTextString(metric))
    w.Write(b)
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
    f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
    if err != nil {
        return err
    }
    n, err := f.Write(data)
    if err == nil && n < len(data) {
        err = io.ErrShortWrite
    }
    if err1 := f.Close(); err == nil {
        err = err1
    }
    return err
}

func Metrics(c *web.C, h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        //reqID := middleware.GetReqID(*c)
        //printStart(reqID, r)

        lw := mutil.WrapWriter(w)

        t1 := time.Now()
        h.ServeHTTP(lw, r)
        //
        if lw.Status() == 0 {
            lw.WriteHeader(http.StatusOK)
        }
        t2 := time.Now()
        f := float64(t2.Sub(t1) / 1000)

        d1 := []byte(fmt.Sprintf("%f\n", f))

        //fp := os.OpenFile("test.log", os.O_WRONLY, FileMod)
        //fmt.Println(f)
        err := WriteFile("/tmp/test.log", d1, 0644)
        if err != nil {
            fmt.Println(err.Error())
        }
        //fp.Write(f+"\n")
        //fp.Close()

        M.h.Observe(f)
        //printEnd(reqID, lw, t2.Sub(t1))
    }

    return http.HandlerFunc(fn)
}