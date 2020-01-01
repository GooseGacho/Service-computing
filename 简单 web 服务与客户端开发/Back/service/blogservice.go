package service 

import (
	"fmt"
	"github.com/cjf/smog/gen-go/rpc"
)

func GetBlogsByTUID(tid, uid int32) (r []*rpc.Blog, err error) {
	fmt.Println("OHH!!")
	//test
	return []*rpc.Blog {
		&rpc.Blog{Bid: 1, Uname: "周星星吖", Iscopy: false, Tname: "Dijkstra 最短路径", Title: "Dijkstra---单源最短路径", Content: testContent(), Readcnt: 8, Cdate: "2019-12-31 15:29:20 ", Ccnt: 2},
	}, nil
}

func GetBlogsByUID(uid int32) (r []*rpc.Blog, err error) {
	//test
	return []*rpc.Blog {
		&rpc.Blog{Bid: 1, Uname: "周星星吖", Iscopy: false, Tname: "Dijkstra 最短路径", Title: "Dijkstra---单源最短路径", Content: testContent(), Readcnt: 8, Cdate: "2019-12-31 15:29:20 ", Ccnt: 2},
	}, nil
}

func GetBlog(bid int32) (r *rpc.Blog, err error) {
	//test
	return 		&rpc.Blog{Bid: 1, Uname: "周星星吖", Iscopy: false, Tname: "Dijkstra 最短路径", Title: "Dijkstra---单源最短路径", Content: testContent(), Readcnt: 8, Cdate: "2019-12-31 15:29:20 ", Ccnt: 2}, nil
}

func testContent() string {
	str := `                             Dijkstra
    [基本思想]:
    总述:从一个最初只含有源点的有向子网开始,逐步扩大到由单源最短路径构成的有向子网为止.
    
    最初只含有源点的有向子网称为入选子网.入选子网以外的顶点组成候选点集.
    
    @ 候选点集的一个顶点被加入到入选子网需满足条件:从源点到该顶点的最短带权路径长度已知,而且除该顶点外,路径的其他顶点都属于入选子网.
    
    1. 开始时,入选子网只含有源点,从源点到源点的最短带权路径长度为0,候选点集的每个顶点都对应这样一条路径:是从源点到该顶点且中间只经过入选子网顶点的路径中最短的一条,这条路径称为候选路径.候选路径集合称为候选路径集.
    
    2. 在候选路径集中选择一条最短的,把终点和相应的边加入到入选子网,入选子网扩大了,候选点集缩小了,这时的候选路径集需要更新,然后再选出最短的,加入到入选子网.
    
    依此类推,直到入选子网顶点集等于网络的顶点集为止.
    
    [实例讲解]:
    
    
    准备:首先将V0加入入选子网,遍集为空集,其他顶点构成候选点集.
    
    (步骤1)将所有V0->其他顶点的路径填写好,无路径写无穷.找出值最小的路径(V0->V2,10),然后将V2加入入选子网.
    
    (步骤2)更新候选路径集.例如V0->V3原来的最短路径长度为无穷,当借助入选子网中的V2后,最短路径更新为V0->V2->V3,长度为(V0,V2)+(V2,V3)=10+50=60.更新时注意:若借助入选子网中的顶点最短路径长度更小就更新.(V0,V1),(V0,V4)(V0,V5)借助V2,最短路径长度没有改变,就不更新.更新后选择最小的(V0,V4,30),将V4加入入选子网.
    
    (步骤3)更新候选路径集.更新时注意:若剩余候选点V1,V3,V5借助入选子网中除源点外的顶点{V2和V4}最短路径可以变得更小就更新最短路径.例如V0->V5原来的最短路径为长度100,借助V4后可以变为(V0,V4)+(V4,V5)=30+60=90,因此更新.最后又是选取最短的路径长度(V0-V4-V3)50,并将V3加入入选子网.
    
    ,,,,,直到所有顶点全部加入到入选子网为止,算法结束
    
    [完整代码]:
    include<iostream>
    include<malloc.h> 
    include<stack>
    define M 100
    define N 100
    using namespace std;
    typedef struct node
    {
      int matrix[N][M];      //邻接矩阵 
      int n;                 //顶点数 
      int e;                 //边数 
    }MGraph; 
    void DijkstraPath(MGraph G,int *dist,int *path,int v0)//v0表示源顶点 
    {
      int i,j,k;
      bool *visited=(bool *)malloc(sizeof(bool)*G.n); 
      //初始化 : 
      for(i=0;i<G.n;i++)    
      {
        if(G.matrix[v0][i]>0&&i!=v0)
        {
          dist[i]=G.matrix[v0][i];//说明V0与i之间有路径,并把最短路径设为原始距离   
          path[i]=v0;     //path记录最短路径上从v0到i需要经过的前一个顶点 
        }
        else
        {
          dist[i]=INT_MAX;    //若i与v0不相邻,则权值置为无穷大 
          path[i]=-1;
        }
        visited[i]=false;//初始化每个节点对应的访问标记.false代表未访问过. 
        path[v0]=v0;
        dist[v0]=0;//源顶点到本身距离为0 
      }
      //V0是源点,从它出发,做访问标记 
      visited[v0]=true;
      for(i=1;i<G.n;i++)//循环扩展n-1次(共n-1个候选点集,逐个加入入选子网) 
      {
        int min=INT_MAX;
        int u;
      //寻找未被访问的且到源点路径长度最小的顶点并做访问标记: 
        for(j=0;j<G.n;j++)    
        {
          if(visited[j]==false&&dist[j]<min)
          {
            min=dist[j];//找到与源点距离最近的点 
            u=j; //用u记录该顶点       
          }
        } 
        visited[u]=true; 
      //更新候选路径集 : 
        for(k=0;k<G.n;k++)
        {
          if(visited[k]==false && G.matrix[u][k]>0 && min+G.matrix[u][k]<dist[k])
          {  //如果经过u到达k比源点直接到达k近,则更新dis[k] 
            dist[k]=min+G.matrix[u][k];
            path[k]=u;//用path[k]记录 v0->k 的前一个点u 
          }
        }        
      }    
    }
     
    void showPath(int *path,int v,int v0)   //打印最短路径上的各个顶点 
    {
      cout<<\"源点 \"<<v0<<"-->"<<\"顶点 \"<<v<<\" 的最短路径为:\";
      stack<int> s;
      int u=v;
      while(u!=v0)
      { //不断把path保存的v-v0的路径 保存到栈中,再逐个弹出就实现了v0-v的路径输出 
        s.push(u);
        u=path[u];
      }
      s.push(u);//此时的v已经与源点v0相等 
      while(!s.empty())
      {
        cout<<s.top()<<" ";
        s.pop();
      }
    } 
     
    int main(int argc, char *arGv[])
    {
      int n,e;     //表示输入的顶点数和边数 
      int i,j;
        int start,end,w;      //表示存在一条边start->end,权值为w
        MGraph G;
        int v0;
        //多组测试案例  
      while(cin>>n>>e && e!=0)
      {   
        int *dist=(int *)malloc(sizeof(int)*n);//dist[i]表示 vo->i 的最短路径长度 
        int *path=(int *)malloc(sizeof(int)*n);//path[v]=u表示v0->v的最短路径中,先经过u再经过v; 
        for(i=0;i<N;i++)
          for(j=0;j<M;j++)
            G.matrix[i][j]=0;
        G.n=n;
        G.e=e;
        for(i=0;i<e;i++)
        {
          cin>>start>>end>>w;
          G.matrix[start][end]=w;
        }
        cin>>v0;        //输入源顶点 
        DijkstraPath(G,dist,path,v0);
        for(i=0;i<n;i++)
        {//打印源点到每一个点的最短路径 
          if(i!=v0)
          {
            showPath(path,i,v0);//打印最短路径上的点 
            cout<<\"  最短路径长度为: \"<<dist[i]<<endl;//及最短路径长度 
          }
        }
      }
      return 0;
	}`
	return str
}