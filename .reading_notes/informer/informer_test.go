package informer

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
	"time"
)

func TestInformer(t *testing.T) {
	config ,err := clientcmd.BuildConfigFromFlags("","/Users/jamlee/.kube/config")
	if err != nil {
		panic(err)
	}
	//通过config 拿到client set客户端
	clientset ,err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	stopch := make(chan struct{})
	defer close(stopch)
	//通过client set客户端以及一个rsync（多久设置一次重新同步 也就是同步间隔时间 如果是0 那么禁用同步功能） 我们拿到一个informer的集合
	sharedInformers := informers.NewSharedInformerFactory(clientset,time.Minute)
	//通过sharedInformers 我们获取到pod的informer
	podInformer := sharedInformers.Core().V1().Pods().Informer()
	//为pod informer添加 controller的 handlerFunc  触发回调函数之后 会通过 addCh 传给 nextCh 管道然后调用controller的对应的handler来做处理
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		//pod资源对象创建的时候出发的回调方法
		AddFunc: func(obj interface{}) {
			obja := obj.(v1.Object)
			fmt.Println(obja)
		},
		//更新回调
		UpdateFunc: func(oldObj, newObj interface{}) {

		},
		//删除回调
		DeleteFunc: func(obj interface{}) {
		},
	})

	//这里会调用reflector的run listandwatch 然后以goroutine的方式运行
	podInformer.Run(stopch)
}

func TestInformerNode(t *testing.T) {
	stopch := make(chan struct{})
	defer close(stopch)

	config , _ := clientcmd.BuildConfigFromFlags("","/Users/jamlee/.kube/config")
	clientset, _ := kubernetes.NewForConfig(config)

	// 直接获取node 列表
	//nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	//fmt.Println(nodes);

	// 通过informer 获取node 列表
	factory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	nodeInformer := factory.Core().V1().Nodes()
	go nodeInformer.Informer().Run(stopch)
	if !cache.WaitForCacheSync(stopch, nodeInformer.Informer().HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	nodes, _ := nodeInformer.Lister().List(labels.NewSelector())
	fmt.Println(nodes);
}
