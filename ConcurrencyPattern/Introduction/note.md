# 并发编程

## 管道(pipe)是一种半双工的通信方式，只能用于父进程与子进程以及同祖先的子进程之间的通信。
    ps aux | grep go

## 信号
> 操作系统信号(signal)是IPC中唯一一种异步的通信方法，它的本质是用软件来模拟硬件的中断机制。

    kill -l
     1) SIGHUP	 2) SIGINT	 3) SIGQUIT	 4) SIGILL	 5) SIGTRAP
     6) SIGABRT	 7) SIGBUS	 8) SIGFPE	 9) SIGKILL	10) SIGUSR1
    11) SIGSEGV	12) SIGUSR2	13) SIGPIPE	14) SIGALRM	15) SIGTERM
    16) SIGSTKFLT	17) SIGCHLD	18) SIGCONT	19) SIGSTOP	20) SIGTSTP
    21) SIGTTIN	22) SIGTTOU	23) SIGURG	24) SIGXCPU	25) SIGXFSZ
    26) SIGVTALRM	27) SIGPROF	28) SIGWINCH	29) SIGIO	30) SIGPWR
    31) SIGSYS	34) SIGRTMIN	35) SIGRTMIN+1	36) SIGRTMIN+2	37) SIGRTMIN+3
    38) SIGRTMIN+4	39) SIGRTMIN+5	40) SIGRTMIN+6	41) SIGRTMIN+7	42) SIGRTMIN+8
    43) SIGRTMIN+9	44) SIGRTMIN+10	45) SIGRTMIN+11	46) SIGRTMIN+12	47) SIGRTMIN+13
    48) SIGRTMIN+14	49) SIGRTMIN+15	50) SIGRTMAX-14	51) SIGRTMAX-13	52) SIGRTMAX-12
    53) SIGRTMAX-11	54) SIGRTMAX-10	55) SIGRTMAX-9	56) SIGRTMAX-8	57) SIGRTMAX-7
    58) SIGRTMAX-6	59) SIGRTMAX-5	60) SIGRTMAX-4	61) SIGRTMAX-3	62) SIGRTMAX-2
    63) SIGRTMAX-1	64) SIGRTMAX	

*其中，编号从1到31的信号属于标准信号（也称为不可靠信号），而编号从34到64的信号属于实时信号（也称为可靠信号）*

- 标准信号只会被记录并处理一次。并且，如果发送给某一个进程的标准信号的种类有多个，那么它们的处理顺序也是完全不确定的。
- 实时信号多个同种类的实时信号都可以记录在案，并且它们可以按照信号的发送顺序被处理。

**信号的来源：键盘键入、硬件故障、系统函数调用和软件中的非法运算<br>
Linux对每一个标准信号都有默认的操作方法：终止进程、忽略该信号、终止进程并保存内存信息、停止进程、恢复进程(若进程已停止)**

Go命令制定了需要被处理的信号

os.Signal接口类型
    
    type Signal interface {
        String() string
        Signal() // to distinguish from other Stringers
    }

syscall.Signal是os.Signal接口的一个实现类型，同时也是一个int类型的别名类型。
>也就是说，每个信号常量都隐含着一个整数值，并且都与它所表示的信号在所属操作系统中的编号一致。

代码包os/signal中的Notify函数用来当操作系统向当前进程发送指定信号时发出通知。
    
    func Notify(c chan<-os.Signal, sig ...os.Signal)
> 参数c是一个发送通道<br>
signal.Notify函数会把当前进程接收到的指定信号放入参数c代表的通道类型值(以下简称signal接收通道)中，这样函数的调用方法就可以从这个signal接收通道中按顺序获取操作系统发来的信号并进行相应的处理了。
<br>第二个参数是一个可变长的参数，这意味着我们在调用signal.Notify函数时，可以在第一个参数值之后再附加任意个os.Signal类型的参数值。

> 除了能够自行处理它们之外，还可以在之后的任意时刻恢复对它们的系统默认操作。
    
    func Stop(c chan<- os.Signal)

函数signal.Stop会取消掉在之前调用signal.Notify 函数时告知signal处理程序需要自行处理若干信号的行为。
1. 只有把当初传递给signal.Notify函数的那个signal接收通道作为调用signal.Stop函数时的参数值，才能如愿以偿地取消掉之前的行为，否则调用signal.Stop函数不会起到任何信号。
2. 这里存在一个副作用，即在之前用于range的for将阻塞，为了消除这种副作用，可以在调用signal.Stop之后，使用内建函数close将该chan关闭。

    