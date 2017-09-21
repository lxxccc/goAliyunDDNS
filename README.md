# goAliyunDDNS
阿里云域名动态解析处理golang版


如果会GO语言，则建议使用源码自己编译，按照自己的喜好修改。   
不会的同学，可以直接使用            
> Windows:   
> goAliyunDDNS_win_x64.exe          
>
> Linux:   
> goAliyunDDNS_linux_x64

将以上提到的文件添加进系统的自动执行策略里，就可以实现动态更新自己的域名到阿里云的解析上了。


## 配置文件说明
需要同步的配置内容在config.json文件内，此文件需要同执行文件放在同一目录下


    AccessKeyID:你的阿里云访问key，在你的阿里云控制台里面的accesskeys里面去找
    AccessKeySecret:同上
    DomainName:你申请的域名，例如xxxx.com
    RR:对应解析配置中的主机记录一项


> 注意：不要使用子用户accessskey，目前阿里的Go语言SDK暂不支持子用户key