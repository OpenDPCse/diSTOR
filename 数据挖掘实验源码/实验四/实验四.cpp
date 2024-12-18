#include<iostream>
#include<string>
#include<fstream>
#include<algorithm>
using namespace std;
////////////////////定义存储结构//////////////////////////////
class Date
{
public:
	string age;
	string income;
	string student;
	string credit;
	string buy;
	void print()
	{
		cout << age<< " " << income << " "  << student  << " " << credit << " "<<buy<<endl;
	}
};

int main()

{  
//////////////////////////读取数据并保存////////////////////////
	char name1[50];
    ifstream infile;
	cout<<"输入要打开的文件：date.txt"<<endl;
	 cin>>name1;
    infile.open(name1,ios::in);  
	if(infile.fail())
	{
		cout << "error open!" << endl;
	}
	
	Date date[100];
	int datesize=0;	
	string iage, iincome,istudent,icredit,ibuy;       //输入的条件
	int y = 0,n = 0,agey = 0,agen = 0;
	int incomey = 0,incomen =0,studenty = 0,studentn = 0,credity = 0,creditn = 0;   //统计出现的次数
	float p1,p2,p3,p4,p5,p6,p7,p8,p9,p10,px1,px2,px3,px4;   //条件概率和类条件概率
	while(!infile.eof())
	{
	infile >> date[datesize].age >> date[datesize].income >> date[datesize].student>> date[datesize].credit>> date[datesize].buy;
	datesize++;
	}
	cout<<"the date are:"<<endl;
	cout << "age "<<"income "<<"student "<<"credit "<<"buy "<<endl;
	for(int i=0;i<datesize;i++)            //输出要处理的数据
	{
		date[i].print();
	}    
	////////////////////////////条件概率///////////////////////////
	for(int j = 0;j<datesize;j++)
	{
	if(date[j].buy=="yes")
	{
	    y++;
	}
	else if(date[j].buy=="no")
	{
		n++;
	}
	}
    p1 = (float)y/(float)datesize;
    p2 = (float)n/(float)datesize;
	cout<<"P(buys_computer = yes) = "<<y<<"/"<<datesize<<"="<<p1<<endl;
    cout<<"P(buys_computer = no) = "<<n<<"/"<<datesize<<"="<<p2<<endl;
	///////////////////类条件概率////////////////////////////
	cout<<"age:"<<endl;
	cin>>iage;
	cout<<"income:"<<endl;
	cin>>iincome;
	cout<<"student:"<<endl;
	cin>>istudent;
	cout<<"credit:"<<endl;
	cin>>icredit;
	for(int k = 0;k<datesize;k++)
	{
	if(date[k].age==iage&&date[k].buy=="yes")
	{
	agey++;
	}
	if(date[k].age==iage&&date[k].buy=="no")
	{
	agen++;
	}
	if(date[k].income==iincome&&date[k].buy=="yes")
	{
	incomey++;
	}
    if(date[k].income==iincome&&date[k].buy=="no")
    {
	incomen++;
	}
	if(date[k].student==istudent&&date[k].buy=="yes")
	{
	studenty++;
	}
	if(date[k].student==istudent&&date[k].buy=="no")
	{
	studentn++;
	}
	if(date[k].credit==icredit&&date[k].buy=="yes")
	{
	credity++;
	}
	if(date[k].credit==icredit&&date[k].buy=="yes")
	{
	creditn++;
	}
		}
	p3=(float)agey/(float)y;
	p4=(float)agen/(float)n;
	p5=(float)incomey/(float)y;
	p6=(float)incomen/(float)n;
	p7=(float)studenty/(float)y;
	p8=(float)studentn/(float)n;
	p9=(float)credity/(float)y;
	p10=(float)creditn/(float)n;
	px1=p3*p5*p7*p9;
	px2=p4*p6*p8*p10;
	px3=px1*p1;
	px4=px2*p2;
	cout<<"P(age = "<<iage<<"|buy = yes = "<<agey<<"/"<<y<<"="<<p3<<endl;
	cout<<"P(age = "<<iage<<"|buy = no = "<<agen<<"/"<<n<<"="<<p4<<endl;
	cout<<"P(income = "<<iincome<<"|buy = yes = "<<incomey<<"/"<<y<<"="<<p5<<endl;
	cout<<"P(income = "<<iincome<<"|buy = no = "<<incomen<<"/"<<n<<"="<<p6<<endl;
	cout<<"P(student = "<<istudent<<"|buy = yes = "<<studenty<<"/"<<y<<"="<<p7<<endl;
	cout<<"P(student = "<<istudent<<"|buy = no = "<<studentn<<"/"<<n<<"="<<p8<<endl;
	cout<<"P(credit = "<<icredit<<"|buy = yes = "<<credity<<"/"<<y<<"="<<p9<<endl;
	cout<<"P(ctedit = "<<icredit<<"|buy = no = "<<creditn<<"/"<<n<<"="<<p10<<endl;

	cout<<"P(X|buy = yes) = "<<px1<<endl;
    cout<<"P(X|buy = no) = "<<px2<<endl;
	cout<<"P(X|buy = yes)P(buy = yes) = "<<px3<<endl;
    cout<<"P(X|buy = no)P(buy = no) = "<<px4<<endl;
	////////////////预测/////////////////////////////////////
	if(px3>px4)
	{
	cout<<"朴素贝叶斯预测buy = yes"<<endl;
	}
	else 
cout<<"朴素贝叶斯预测buy =no"<<endl;
 infile.close();//关闭文件
system( "PAUSE ");


}
