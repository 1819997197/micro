# ch11 docker简介

2010年，几个搞IT的年轻人，在美国旧金山成立了一家名叫“dotCloud”的公司。

这家公司主要提供基于PaaS的云计算技术服务。具体来说，是和LXC有关的容器技术。

![Image text](https://github.com/1819997197/micro/blob/master/ch11/picture/lxc.jpg)

后来，dotCloud公司将自己的容器技术进行了简化和标准化，并命名为——Docker。

![Image text](https://github.com/1819997197/micro/blob/master/ch11/picture/docker.jpg)




### docker

我们具体来看看Docker，**Docker本身并不是容器**，它是创建容器的工具，是应用容器引擎。

想要搞懂Docker，其实看它的两句口号就行。

第一句，是** “Build, Ship and Run”**。

![Image text](https://github.com/1819997197/micro/blob/master/ch11/picture/build_ship_run.jpg)

也就是，“搭建、发送、运行”，三板斧。

举个例子：

我来到一片空地，想建个房子，于是我搬石头、砍木头、画图纸，一顿操作，终于把这个房子盖好了。

![Image text](https://github.com/1819997197/micro/blob/master/ch11/picture/home.jpg)

结果，我住了一段时间，想搬到另一片空地去。这时候，按以往的办法，我只能再次搬石头、砍木头、画图纸、盖房子。

但是，跑来一个老巫婆，教会我一种魔法。

这种魔法，可以把我盖好的房子复制一份，做成“镜像”，放在我的背包里。

![Image text](https://github.com/1819997197/micro/blob/master/ch11/picture/home2depository.jpg)

等我到了另一片空地，就用这个“镜像”，复制一套房子，摆在那边，拎包入住。

![Image text](https://github.com/1819997197/micro/blob/master/ch11/picture/depository2home.jpg)

怎么样？是不是很神奇？

所以，Docker的第二句口号就是：**“Build once，Run anywhere（搭建一次，到处能用）”**。

Docker技术的三大核心概念，分别是：
> * 镜像（Image）
> * 容器（Container）
> * 仓库（Repository）

我刚才例子里面，**那个放在包里的“镜像”，就是Docker镜像。而我的背包，就是Docker仓库。我在空地上，用魔法造好的房子，就是一个Docker容器。**

说白了，这个Docker镜像，是一个特殊的文件系统。它除了提供容器运行时所需的程序、库、资源、配置等文件外，还包含了一些为运行时准备的一些配置参数（例如环境变量）。镜像不包含任何动态数据，其内容在构建之后也不会被改变。

也就是说，每次变出房子，房子是一样的，但生活用品之类的，都是不管的。谁住谁负责添置。

每一个镜像可以变出一种房子。那么，我可以有多个镜像呀！

也就是说，我盖了一个欧式别墅，生成了镜像。另一个哥们可能盖了一个中国四合院，也生成了镜像。还有哥们，盖了一个非洲茅草屋，也生成了镜像。。。

这么一来，我们可以交换镜像，你用我的，我用你的，岂不是很爽？

![Image text](https://github.com/1819997197/micro/blob/master/ch11/picture/villa.jpg)

于是乎，就变成了一个大的公共仓库。

负责对Docker镜像进行管理的，是**Docker Registry服务**（类似仓库管理员）。

不是任何人建的任何镜像都是合法的。万一有人盖了一个有问题的房子呢？

所以，Docker Registry服务对镜像的管理是非常严格的。

最常使用的Registry公开服务，是官方的**Docker Hub**，这也是默认的 Registry，并拥有大量的高质量的官方镜像。


### 参考资料
```
https://zhuanlan.zhihu.com/p/53260098
https://yeasy.gitbooks.io/docker_practice/content/
```