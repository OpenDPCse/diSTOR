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
	char date[10];
	int sn;
    int id;
	float num;
	float price;
};
int main()
{  
	char name1[50],name2[50];
    ifstream infile;
	cout<<"ѡ��Ҫ�򿪵��ļ���1019n.txt 1020n.txt 1021n.txt"<<endl;
	cin>>name1;
    infile.open(name1,ios::in);  
	/*string contents;*/
	if(infile.fail())
	{
		cout << "error open!111" << endl;
	}
	cout<<"Ҫ������ļ�����"<<endl;
 	cin>>name2;
	ofstream outfile(name2,ios::out);
	if(!outfile){
     	cout<<"open eror!222"<<endl;
		exit(1);
 	}
	
	Sales_n sal[10000];
	int sal_size=0;
	float total[10][5]={0};
	
	while(!infile.eof()){//���ļ� 
		infile >> sal[sal_size].serial >> sal[sal_size].market >> sal[sal_size].date>> sal[sal_size].sn>> sal[sal_size].id>> sal[sal_size].num>> sal[sal_size].price;
		sal_size++;
	}
	
	
	for(int i=0;i<sal_size;i++){
       	char p = sal[i].date[7];
		if(sal[i].id/100==10010 ){
			switch(p){
				case '3':total[0][0]+=sal[i].num*sal[i].price;break;
		        case '4':total[1][0]+=sal[i].num*sal[i].price;break;
				case '5':total[2][0]+=sal[i].num*sal[i].price;break;
				case '6':total[3][0]+=sal[i].num*sal[i].price;break;
				case '7':total[4][0]+=sal[i].num*sal[i].price;break;
				case '8':total[5][0]+=sal[i].num*sal[i].price;break;
				case '9':total[6][0]+=sal[i].num*sal[i].price;break;
			}
		}
		if(sal[i].id/100==10020 ){
			switch(p){
				case '3':total[0][1]+=sal[i].num*sal[i].price;break;
		        case '4':total[1][1]+=sal[i].num*sal[i].price;break;
				case '5':total[2][1]+=sal[i].num*sal[i].price;break;
				case '6':total[3][1]+=sal[i].num*sal[i].price;break;
				case '7':total[4][1]+=sal[i].num*sal[i].price;break;
				case '8':total[5][1]+=sal[i].num*sal[i].price;break;
				case '9':total[6][1]+=sal[i].num*sal[i].price;break;
			}
		}
		if(sal[i].id/100==10030){
			switch(p){
				case '3':total[0][2]+=sal[i].num*sal[i].price;break;
		        case '4':total[1][2]+=sal[i].num*sal[i].price;break;
				case '5':total[2][2]+=sal[i].num*sal[i].price;break;
				case '6':total[3][2]+=sal[i].num*sal[i].price;break;
				case '7':total[4][2]+=sal[i].num*sal[i].price;break;
				case '8':total[5][2]+=sal[i].num*sal[i].price;break;
				case '9':total[6][2]+=sal[i].num*sal[i].price;break;
			}
		}
		else if(sal[i].id/100==10088){
			switch(p){
				case '3':total[0][3]+=sal[i].num*sal[i].price;break;
		        case '4':total[1][3]+=sal[i].num*sal[i].price;break;
				case '5':total[2][3]+=sal[i].num*sal[i].price;break;
				case '6':total[3][3]+=sal[i].num*sal[i].price;break;
				case '7':total[4][3]+=sal[i].num*sal[i].price;break;
				case '8':total[5][3]+=sal[i].num*sal[i].price;break;
				case '9':total[6][3]+=sal[i].num*sal[i].price;break;
			}
		}
	}

	if (outfile){
		cout<<"����"<<'\t'<<'\t'<<"10010�� "<<"10020����Ʒ "<<"10030�׺ͷ� "<<"10088��������Ʒ "<<endl;
		int j = 20030413;
        for (int i=0;i<7;++i){
                outfile<<" "<< total[i][0]<<'\t'<<total[i][1]<<'\t'<<total[i][2]<<'\t'<<total[i][3]<<'\t'<<endl;
				cout<<j<<'\t'<< total[i][0]<<'\t'<<total[i][1]<<'\t'<<total[i][2]<<'\t'<<total[i][3]<<'\t'<<endl;
				j++;
        }
    }
    else
        cerr<<"�޷����ļ�333��"<<endl;

	ifstream infile2(name2,ios::in);
	
	int m=0;
	while(!infile.eof()){
		infile >> total[m][0] >> total[m][1] >>total[m][2]>>total[m][3];
	 	 m++;
	}
	if(infile2.fail()){
		cout << "error open!" << endl;
	}
	float sum=0.0;
	for(int i=0;i<7;++i){
		sum+=total[i][2];
	}
	cout<<"2020������10010������Ʒ13�����۶�Ϊ��"<<total[0][0]<<endl;
	cout<<"2020������10030�׺ͷ�����Ʒ�����ܶ�Ϊ��"<<sum<<endl;
	
 	infile.close();//�ر��ļ�
 	infile2.close();//�ر��ļ�
 	outfile.close();//�ر��ļ�
	system( "PAUSE ");
}
