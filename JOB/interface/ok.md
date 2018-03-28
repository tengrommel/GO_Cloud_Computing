#time

# 接口
>为什么需要接口

归一化处理

定义接口实现规范

    type Writer inteface {
        Write(p []byte) (n int, err error)
    }

