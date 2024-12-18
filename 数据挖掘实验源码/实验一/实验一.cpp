#include<iostream>
#include<string>
#include<fstream>
#include<algorithm>
using namespace std;
class Sales_n
{
public:
	string serial;
	int market;
	int posno;
	string date;
	int sn;
    int id;
	float num;
	float price;
	float total;
	void print(){
		cout << serial<< " " << market << " "	<< posno	<< " " << date << " "<<sn<<" "<<id<<" "<<num<<" "<<price<<" "<<total<<endl;
	}
};

bool cmp(Sales_n sale1,Sales_n sale2){//排序规则 
	if(sale1.sn!=sale2.sn){
		return sale1.sn<sale2.sn;
	}else{
		return sale1.id<sale2.id;
	}
}

int main()
{  
	char name1[50],name2[50];
    ifstream infile;
	cout<<"选择要打开的文件：1019.txt 1020.txt 1021.txt"<<endl;
	cin>>name1;
    infile.open(name1,ios::in);  
	if(infile.fail())
	{
		cout << "error open!" << endl;
	}
	cout<<"要保存的文件名："<<endl;
 	cin>>name2;
	ofstream outfile(name2,ios::out);
	if(!outfile){
     cout<<"open eror!"<<endl;
     exit(1);
 }
	Sales_n sal[10000];
	int sal_size=0;
	while(!infile.eof())
	{
	infile >> sal[sal_size].serial >> sal[sal_size].market >>sal[sal_size].posno>> sal[sal_size].date>> sal[sal_size].sn>> sal[sal_size].id>> sal[sal_size].num>> sal[sal_size].price>>sal[sal_size].total;
	sal_size++;
	}
	cout<<"文档"<<name1<<"的长度是"<<sal_size<<endl;
	int l=0;
	sort(sal,sal+sal_size,cmp);//进行排序 
	for (int i =0;i<sal_size;i++)
	{  
		if(sal[i].num<0){//购买数量不能为负值,取反 
			sal[i].num=-sal[i].num;
		}
		sal[i].date.assign(sal[i].serial,0,8);//设置date为流水号的前八位 
		//去掉总额 
		//去掉pos机号码 
		outfile<<sal[i].serial<<"\t"<<sal[i].market<<"\t"<<sal[i].date<<"\t"<<sal[i].sn<<"\t"
			<<sal[i].id<<"\t"<<sal[i].num<<"\t"<<sal[i].price<<endl;
			
		l=i;
	}
	cout<<"预处理成功！"<<endl;
  
 	infile.close();//关闭文件
 	outfile.close();//关闭文件
	system( "PAUSE ");
	return 0;
}
