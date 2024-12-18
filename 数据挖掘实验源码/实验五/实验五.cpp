#include<iostream>
#include<string>
#include<fstream>
#include<algorithm>
#include <math.h>
using namespace std;

#define k 3          //聚类数
#define n 2        //数据维数
#define size 30     //数据大小
////////////////////定义存储结构//////////////////////////////
typedef struct {
    double d[n];
    double distance[k];
}Data;
typedef struct {
    Data center[k];  // 即簇类中心 
	int cluster[k][size]; //簇数组
	int cluster_num[k];// 簇类中一组数据的编号
    Data old_center[k];
	int isend[k][n]; //各簇类中心是否相等标示值
	int is;   
}Tdata;

////////声明函数//////////////////////////////

void input_data();
void Init_center();
void calculate_distance();
double Euclid(int x, int y);
void new_cluster();
void new_center();
void output_info();
void comp();

Data data[size];
Tdata td;

///////////////读入数据////////////////////////
void Init()
{
char name1[50];
    ifstream infile;
	cout<<"输入要打开的文件：data.txt"<<endl;
	 cin>>name1;
    infile.open(name1,ios::in);  
	if(infile.fail())
	{
		cout << "error open!" << endl;
	}
	for(int i = 0;i < size;i++)
		for(int j = 0;j<n;j++)
		{
		infile>>data[i].d[j];
		}
	
	cout<<"the data are:"<<endl;
	for(int i=0;i<size;i++) {           //输出要处理的数据
	for(int j = 0;j<n;j++)
	{
		cout<<data[i].d[j]<<"  ";
	} 
	cout<<endl;
	}
	 infile.close();//关闭文件
}
//////////////初始化质心//////////////////////////////
void Init_center()
{	
	for(int i = 0;i<k;i++){
		cout<<"初始质心"<<i+1<<":"<<endl;
		for(int j = 0;j<n;j++)
	{
     td.center[i].d[j] = data[i].d[j];
	  cout<<td.center[i].d[j]<<"  "; 
	}
cout<<endl;
	}	
}


  ///////计算数据到K个质心的欧几里德距离//////////

void calculate_distance()
{
    int i, j;
    for(i = 0; i < size; i++)
        for(j = 0; j < k; j++){
            data[i].distance[j] = Euclid(i, j);   //i表示第几个数组j表示距离第几个质心
        }
}

//////////计算一组数组到质心的欧几里得距离//////////////
double Euclid(int x, int y)
{
    double distance = 0;
    for(int i = 0; i < n; i++){
        distance += pow((data[x].d[i] - td.center[y].d[i]), 2);
    }
    distance = sqrt(distance);
    return distance;
}
//////////将数据进行簇归类///////////////////
void new_cluster()
{
    int i, j;
    double min;
    for(i = 0; i < k; i++)      //初始化编号
        td.cluster_num[i] = 0;
    for(i = 0; i < size; i++){ 
        int index = 0;              //找出最小的欧几里德距离编号
        min = data[i].distance[0];
        for(j = 1; j < k; j++){     // 筛选到簇心欧几里德最小的值
            if(data[i].distance[j] < min){
                min = data[i].distance[j];
                index = j;
            }
        }
        //划分簇集
        td.cluster[index][td.cluster_num[index]++] = i;   
    }
}
///////////更新质心///////////////////////////////
void new_center()
{
    int i, j, m;
    double sum;
	

    for(i = 0; i < k; i++)
        for(j = 0; j < n; j++){
            sum = 0;
			td.old_center[i].d[j] =  td.center[i].d[j];
            for(m = 0; m < td.cluster_num[i]; m++){
               // 第i个簇的第j维数的所有数据和
                sum += data[td.cluster[i][m]].d[j];
            }
           // 取平均数得到新的簇中心
            td.center[i].d[j] = sum / td.cluster_num[i];
        }
}

////////比较质心////////////////////////

void comp( )
 {
	 int i , j,m = 0;
	 new_center();
	 for(i = 0;i<k;i++){
		 for(j=0;j<n;j++){
            
			 if(td.old_center[i].d[j] !=  td.center[i].d[j]) td.isend [i][n]= 0;
			 else td.isend [i][n]= 1;
		 }
	 }
	for(i = 0;i<k;i++){
		 for(j=0;j<n;j++){
			 td.is=td.isend [i][n]*td.is;
		 }
	}
	if(td.is==0)td.is=1;
	else td.is=0;
 
 }
////输出结果函数/////////////////

void output_info( )
{
    int i, j, m;

    for(i = 0; i < k; i++){ 
       
		cout<<"质心"<<i+1<<":"<<endl;
            for(m = 0; m < n; m++)
                cout<<td.center[i].d[m]<<"  ";
			cout<<endl;

			cout<<"簇类"<<i+1<<":"<<endl;
        for(j = 0; j < td.cluster_num[i]; j++){            
            for(m = 0; m < n; m++){
                
                  cout<<data[td.cluster[i][j]].d[m]<<"  ";
            }
            cout<<endl;
        }
    }
}

int main ()
{
	int count = 0;
Init();
Init_center();
td.is =1;
 while (td.is)
{
 calculate_distance(); 
 new_cluster(); 

 count ++;
cout<<"______________________________________________________"<<endl;
cout<<"第"<<count<<"次聚类："<<endl;
output_info();
comp(); //比较质心
		}
 
system( "PAUSE ");
return 0;
}
