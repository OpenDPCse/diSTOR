#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <map>
#include <cmath>
#include <bitset>
#include <algorithm>
#include <iomanip>
using namespace std;
const int minsup=2;  //������С֧�ֶ�
map<string,int> items_count;   //ͳ�Ƹ��������Ŀ
vector<string> mergeItem(vector<string> vect1,vector<string> vect2,int round);     //�ϲ������µĺ�ѡ�
int isExist(vector<string> item,vector<vector<string> >items);         //�ж��item�Ƿ��Ѿ����ں�ѡ�����items��,�����򷵻�1
vector<string> mergeItem(vector<string> vect1,vector<string> vect2,int round)            //�ж�������Ƿ���Ժϲ�(Ҫ��ֻ��һ�ͬ)��һ���µ������Ϊ��ѡ����
{


   ////////////////////////////////////////////��֦����////
   int count=0;     //ͳ������vector����ͬ�������Ŀ
   vector<string> vect;
   map<string,int> tempMap;   //�����ж�����vector���ظ�����
   for(vector<string>::size_type st=0;st<vect1.size();st++)
   {
      tempMap[vect1[st]]++;
      vect.push_back(vect1[st]);
   }
   for(int st=0;st<vect2.size();st++)
   {
      tempMap[vect2[st]]++;
      if(tempMap[vect2[st]]==2)  //��ʾ��������ͬ
      {
         count++;
      }
      else
      {
         vect.push_back(vect2[st]);
      }
   }
   if((count+1)!=round)       //Ҫ��������Ŀ��ֻ��һ����Ŀ����ͬ����������ͬ,��:I1 I2 I4 ��I1 I2 I3
   {
      vect.clear();
   }
   return vect;
}




int isExist(vector<string> item,vector<vector<string> >items)  //�ж��item�Ƿ��Ѿ����ں�ѡ�����items��,�����򷵻�1
{
   ////////////////�ظ��ı���:I1I2I3��I2I3I1
    int count;   //ͳ����ͬ�������Ŀ
    if(!items.empty())
    {
       for(vector<vector<string> >::size_type ix=0;ix!=items.size();ix++)
       {
           count=0;
           for(vector<string>::size_type iy=0;iy!=items[ix].size();iy++)
           {
               for(vector<string>::size_type iz=0;iz!=item.size();iz++)
               {
                   if(item[iz]==items[ix].at(iy))
                   {
                      count++;
                   }
               }
           }
           if(count==item.size())     //��ʾ����
           {
              return 1;
           }
       }
     }
    return 0;
}
int main()
{
  vector<vector<string> > datavec;        //ԭʼ�����
  vector<vector<string> > candidatevec;   //��ѡ�
  vector<vector<string> > frequentvec;   //Ƶ���
  vector<map<string,int> > bitmap;     //�ж�ĳ����Ŀ��ĳһ���������Ƿ���ڣ�������ֵΪ1����֮Ϊ0
  long trancount=0;   //ԭʼ��������
  
	char name1[50];
    ifstream file;
	cout<<"ѡ��Ҫ�򿪵��ļ���new1.txt new2.txt new3.txt"<<endl;
	 cin>>name1;
    file.open(name1,ios::in);  
	/*string contents;*/
	
  if(!file)        //����ļ��Ƿ�򿪳ɹ�
  {
     cout<<"Fail to open data file!"<<endl;
     return 1;
  }
  else
  {
     string temp;
     vector<string> item;   //�����ʱvector 
     int begin,end;
     while(getline(file,temp))     //һ��һ�ж�������
     {
        trancount++;
        begin=0; 
		 temp.erase(0,temp.find_first_not_of("\r\t\n "));   //ȥ���ַ����ײ��Ŀո�
       temp.erase(temp.find_last_not_of("\r\t\n")+1); 
        while((end=temp.find('\t',begin))!=string::npos)    //ÿһ�������е�������'\t'Ϊ�ָ�����
        {
           item.push_back(temp.substr(begin,end-begin));   //��ÿһ�������item��
           begin=end+1;
        }
        item.push_back(temp.substr(begin));     //һ�������е����һ��
        datavec.push_back(item);       //��һ�������е��������һ�����������һ�����vector��
        item.clear();   //���item
     }
     cout<<"Press Enter to continue the processing";  //pause
     getchar();
     map<string,int> item_map; 
     for(vector<vector<string> >::size_type ix=0;ix!=datavec.size();++ix)
     {
        for(vector<string>::size_type iy=0;iy!=datavec[ix].size();++iy)
        {
         items_count[datavec[ix].at(iy)]++;    //����ļ�����1
         item_map[datavec[ix].at(iy)]=1;       //��ʾ����Ŀ�ڸ������д��ڣ�ֵΪ1������Ĭ��Ϊ0
        }
        bitmap.push_back(item_map);
        item_map.clear();      //����һ��Ҫ���һ��
     }
     map<string,int>::const_iterator map_it=items_count.begin();
     cout<<"��ѡ1�:"<<endl;
     while(map_it!=items_count.end())      //�����ѡ1�
     {
        cout<<"{"<<map_it->first<<"}"<<endl;
        map_it++;
     }
     cout<<"Press Enter to continue the processing";  //pause
     getchar();
     map_it=items_count.begin();
     cout<<"Ƶ��1�(minsup=2):"<<endl;
     while(map_it!=items_count.end())          //Ƶ��1�
     {
        if(map_it->second>minsup)    //֧�ֶȴ���2
        {
            cout.setf(ios::fixed);
			 cout<<"{"<<map_it->first<<"}"<<" ֧�ֶ�:"<<setprecision(6)<<map_it->second<<endl;
            item.push_back(map_it->first);
            frequentvec.push_back(item);   //�����ѡ1���vector��
            item.clear();    
        }
        map_it++;
     }
     if(!frequentvec.empty())   //�ж�Ƶ��1��Ƿ�Ϊ�գ�Ϊ�����˳�
     {
         cout<<"Press Enter to continue the processing";  //pause
         getchar();
         int round=1;     //���ɺ�ѡ��ִ�
         int found;    //�Ƿ�����з�Ƶ�����Ӽ�,Ϊ1��ʾ���У��еĻ����м�֦�������I1I4Ϊ��Ƶ�������I1I2I4Ҫ��֦��
         string tempstr;
         vector<string> tempvec;
         do
         {
            //������һ�ֵĺ�ѡ�
            vector<vector<string> >::size_type st=frequentvec.size();
            candidatevec.clear();         //�����һ�ֵĺ�ѡ�
            for(vector<vector<string> >::size_type st1=0;st1<st;st1++)
            {
                for(vector<vector<string> >::size_type st2=st1+1;st2<st;st2++)
                {
                    found=0;
                    item=mergeItem(frequentvec[st1],frequentvec[st2],round);    //���ú����ϲ�������һ�ֵĺ�ѡ�
                    if(!item.empty()&&!isExist(item,candidatevec))   //�������жϴ���󷵻ص�vector��Ϊ���һ������ڸ��������Ϊ��ѡ������ѡvector��
                    {
                        ////////ʵ�ּ�֦//////////////////////////
                        string teststr;
                        int testint;
                        tempvec=item;
                        sort(tempvec.begin(),tempvec.end());
                        while(next_permutation(tempvec.begin(),tempvec.end()))   //�������е����I1I2I4,Ҫ���I1I4I2��������I2I1I4�����ж�������I1I4�����Ƶ���
                        {
                           for(vector<string>::size_type tempst=0;tempst!=tempvec.size();tempst++) //ƴ�ӳ����ַ������
                           {
                                tempstr+=tempvec[tempst];                                
                           }
                           for(map<string,int>::const_iterator tempit=items_count.begin();tempit!=items_count.end();tempit++)
                           {
                                if(tempit->second<minsup)  //��Ƶ���
                                {
                                    if(tempstr.find(tempit->first)!=string::npos)   //��ʾ�����з�Ƶ�����
                                   {
                                       found=1;
                                       teststr=tempit->first;
                                       testint=tempit->second;
                                       break;
                                   }
                                }
                           }
                           tempstr.erase();
                           if(found)   //������Ƶ�����
                           {
                              break;
                           }
                           
                        }
                        if(!found)     //ֻ�в������з�Ƶ������ż����ѡ��У������֦��
                           candidatevec.push_back(item);
                       /* else
                        {
                           cout<<"��ȥ�:";
                           for(vector<string>::size_type st2=0;st2!=item.size();st2++)
                             //  cout<<item[st2];
                           //cout<<" ���з�Ƶ�����:"<<teststr<<" "<<testint<<"/"<<trancount<<"="<<((float)(testint)/(float)trancount);
                           //cout<<endl;
                        }*/
                        found=0;   //����
                    }
                    
                }
            }
            frequentvec.clear();         //�����һ�ֵ�Ƶ���
            round++;
            cout<<"��ѡ"<<round<<"�:"<<endl;
            for(vector<vector<string> >::size_type ix=0;ix!=candidatevec.size();++ix)      //�����ѡ�
            {
               cout<<"{";
               for(vector<string>::size_type iy=0;iy!=candidatevec[ix].size();++iy)
               {
                cout<<candidatevec[ix].at(iy)<<' ';
               }
               cout<<"}"<<endl;
            }
            if(candidatevec.empty())  //��ѡ�Ϊ��
            {
               cout<<"��ѡ"<<round<<"�Ϊ��!"<<endl;
            }
            int flag;    //���ĳ�����ĳ���������Ƿ���֣�����Ϊ1��������Ϊ0���磺{I1I2}
            int count;   //ͳ��ĳ���뼯���������׵������г��ֵĴ���
            string tempstr;   //��ʱstring�����ڴ��Ӹ������һ���ַ����� ��: I1 I2  I3  ����Ϊ"I1I2I3"
            int mark;    //Ϊ����ִ�ж�����ַ������ӹ���
            for(vector<vector<string> >::size_type sx=0;sx!=candidatevec.size();++sx)      //������һ�ֵ�Ƶ���
            {
                mark=1;
                count=0;
                for(vector<map<string,int> >::size_type sy=0;sy!=bitmap.size();++sy)
                {
                    flag=1;       //��ʼ��Ϊ1�������
                    for(vector<string>::size_type sz=0;sz!=candidatevec[sx].size();++sz)
                    {
                       if(bitmap[sy][candidatevec[sx].at(sz)]==0)   //����ĳһ��������ڣ���û�����
                       {
                           flag=0;
                       }
                       if(mark==1)   //ֻ����һ�Σ���I1I2    ����Ϊ10��I1I2�Ĵ���
                       {
                           tempstr+=candidatevec[sx].at(sz);  //�����ַ���
                       }
                    }
                    
                    if(flag)  //flag��ȻΪ1����ʾ����ڸ��������г����ˣ�������1
                    {
                       count++;
                    }
                    mark++;
                }
                
                if(count>minsup)   //֧�ֶȴ���2
                {
                    frequentvec.push_back(candidatevec[sx]);        //����Ƶ���
                }
                items_count[tempstr]=count;       //��Ӧ����ļ���ֵ
                /////////�����ʱ���ɵ�tempstrΪI1I2I3,Ϊ���ں���������Ŷȵļ��㣬������Ҫ����I2I1I3,I1I3I2����ϣ���
                //��items_count�и����Ǹ����I1I2I3��ͬ��ֵ
                sort(candidatevec[sx].begin(),candidatevec[sx].end());    //����
                string tempstr2;
                while(next_permutation(candidatevec[sx].begin(),candidatevec[sx].end()))  //ȡ��һ�������
                {
                     for(vector<string>::size_type tempst=0;tempst!=candidatevec[sx].size();tempst++) //ƴ�ӳ����ַ������
                     {
                         tempstr2+=candidatevec[sx][tempst];
                     }
                     items_count[tempstr2]=count;  //��Ӧ����ļ���ֵ
                     tempstr2.erase();
                }               
                tempstr.erase();
            }
            cout<<"Press Enter to continue the processing";  //pause
            getchar();
            if(!frequentvec.empty())     //Ƶ�����Ϊ��
            {
                 
                 cout<<"Ƶ��"<<round<<"�(minsup=2):"<<endl;
                 for(int sx=0;sx!=frequentvec.size();++sx)      //���Ƶ���
                 {
                     cout.setf(ios::fixed);
                     cout<<"{"; 
                     for(vector<string>::size_type sz=0;sz!=frequentvec[sx].size();++sz)
                     {
                        cout<<frequentvec[sx].at(sz)<<' ';
                        tempstr+=frequentvec[sx].at(sz);  //�����ַ���
                     }
                     cout<<"}";

    
                     cout<<endl;
                     tempstr.erase();
                 }
                 cout<<"Press Enter to continue the processing";  //pause
                 getchar();
             }
             else
             {
                 cout<<"û��"<<round<<"-Ƶ���,Apriori�㷨����!"<<endl;
             }
         }while(!frequentvec.empty());   //Ƶ�����Ϊ��,��ѭ������
         
        
         file.close();
      
         return 0;
     }
     else
     {         
         return 0;
     }    //end of if(!frequentvec.empty())
     
  }//end of if(!file)
}

