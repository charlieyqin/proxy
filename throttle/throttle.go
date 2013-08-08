package throttle

type Throttler interface {
    Acquire() error
    Release() error
}

type CountingThrottler struct {
    sema chan struct {}
}

func (c *CountingThrottler) Acquire () error {
    <-c.sema
    return nil
}

func (c *CountingThrottler) Release () error {
    var v struct {}
    c.sema <- v
    return nil
}

func NewCountingThrottler(maxConn uint64) Throttler {
    ct := CountingThrottler{make(chan struct{}, maxConn)}
    for i := uint64(0); i < maxConn; i++ {
        var v struct {}
        ct.sema <- v
    }
    return ct
}

