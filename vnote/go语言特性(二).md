# go语言特性(二)

## 线程枷锁

### 锁

Go 语言不仅仅提供基于 CSP 的通讯模型，也支持基于共享内存的多线程数据访问

Sync 包提供了锁的基本原语

- sync.Mutex 互斥锁
    - Lock() 加锁
    - Unlock() 解锁
    ```go
    // Kubernetes 中的 informer factory
    // Start initializes all requested informers.
    func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
        f.lock.Lock()
        defer f.lock.Unlock()
        for informerType, informer := range f.informers {
            if !f.startedInformers[informerType] {
                go informer.Run(stopCh)
                f.startedInformers[informerType] = true
            }
        }
    }
    ```
- sync.RWMutex 读写分离锁
    - 不限制并发读，只限制并发写和并发读写
- sync.WaitGroup
    - 等待一组 goroutine 返回
    ```go
    // CreateBatch create a batch of pods. All pods are created before waiting.
    func (c *PodClient) CreateBatch(pods []*v1.Pod) []*v1.Pod {
        ps := make([]*v1.Pod, len(pods))
        var wg sync.WaitGroup
        for i, pod := range pods {
            wg.Add(1)
            go func(i int, pod *v1.Pod) {
                defer wg.Done()
                defer GinkgoRecover()
                ps[i] = c.CreateSync(pod)
            }(i, pod)
        }
        wg.Wait()
        return ps
    }
    ```
- sync.Once 
    - 保证某段代码只执行一次
- sync.Cond
    - 让一组 goroutine 在满足特定条件时被唤醒
    ```go
    //Kubernetes 中的队列，标准的生产者消费者模式
    // cond: sync.NewCond(&sync.Mutex{}),
    // Add marks item as needing processing.
    func (q *Type) Add(item interface{}) {
        q.cond.L.Lock()
        defer q.cond.L.Unlock()
        if q.shuttingDown {
            return
        }
        if q.dirty.has(item) {
            return
        }
        q.metrics.add(item)
        q.dirty.insert(item)
        if q.processing.has(item) {
            return
        }
        q.queue = append(q.queue, item)
        q.cond.Signal()
    }
    // Get blocks until it can return an item to be processed. If shutdown = true,
    // the caller should end their goroutine. You must call Done with item when you
    // have finished processing it.
    func (q *Type) Get() (item interface{}, shutdown bool) {
        q.cond.L.Lock()
        defer q.cond.L.Unlock()
        for len(q.queue) == 0 && !q.shuttingDown {
            q.cond.Wait()
        }
        if len(q.queue) == 0 {
            // We must be shutting down.
            return nil, true
        }
        item, q.queue = q.queue[0], q.queue[1:]
        q.metrics.get(item)
        q.processing.insert(item)
        q.dirty.delete(item)
        return item, false
    }
    ```

