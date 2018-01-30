# js

    // 闭包外部参数赋值内部执行
    for (var i = 0; i < 10; i++) {
        (function(i) {
            setTimeout(function() {
                console.log(i);
            }, 1000)
        })(i)
    }
    
    // 闭包内部参数赋值返回执行函数
    for (var i = 1; i <= 5; i++) {
        setTimeout((function(i) {
            return function() {
                console.log(i);
            }
        })(i), 1000);
    }
    
# GO
    package main
    
    import (
    	"fmt"
    	"sync"
    	"time"
    )
    var wg sync.WaitGroup
    
    func main() {
    	// 非规律性执行一般都会输入5
    	wg.Add(10)
    	for i:=0;i<5;i++{
    		go func() {
    			fmt.Println(i)
    			defer wg.Done()
    		}()
    	}
    
    	time.Sleep(time.Second)
    	// 规律输入
    	for i:=0;i<5;i++{
    		go func(i int) {
    			fmt.Println(i)
    			defer wg.Done()
    		}(i)
    	}
    	wg.Wait()
    }