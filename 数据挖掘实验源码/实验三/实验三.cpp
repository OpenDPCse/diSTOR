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
const int minsup=2;  //设置最小支持度
map<string,int> items_count;   //统计各个项集的数目
vector<string> mergeItem(vector<string> vect1,vector<string> vect2,int round);     //合并生成新的候选项集
int isExist(vector<string> item,vector<vector<string> >items);         //判断项集item是否已经存在候选项集集合items中,存在则返回1
vector<string> mergeItem(vector<string> vect1,vector<string> vect2,int round)            //判断两个项集是否可以合并(要求只有一项不同)成一个新的项集（做为候选集）
{


   ////////////////////////////////////////////剪枝工作////
   int count=0;     //统计两个vector中相同的项的数目
   vector<string> vect;
   map<string,int> tempMap;   //辅助判断两个vector中重复的项
   for(vector<string>::size_type st=0;st<vect1.size();st++)
   {
      tempMap[vect1[st]]++;
      vect.push_back(vect1[st]);
   }
   for(int st=0;st<vect2.size();st++)
   {
      tempMap[vect2[st]]++;
      if(tempMap[vect2[st]]==2)  //表示这两项相同
      {
         count++;
      }
      else
      {
         vect.push_back(vect2[st]);
      }
   }
   if((count+1)!=round)       //要求两个项目集只有一个项目不相同，其他都相同,如:I1 I2 I4 和I1 I2 I3
   {
      vect.clear();
   }
   return vect;
}




int isExist(vector<string> item,vector<vector<string> >items)  //判断项集item是否已经存在候选项集集合items中,存在则返回1
{
   ////////////////重复的比如:I1I2I3和I2I3I1
    int count;   //统计相同的项的数目
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
           if(count==item.size())     //表示存在
           {
              return 1;
           }
       }
     }
    return 0;
}
int main()
{
  vector<vector<string> > datavec;        //原始数据项集
  vector<vector<string> > candidatevec;   //候选项集
  vector<vector<string> > frequentvec;   //频繁项集
  vector<map<string,int> > bitmap;     //判断某个项目在某一个事务中是否存在，存在则值为1，反之为0
  long trancount=0;   //原始事务总数
  
	char name1[50];
    ifstream file;
	cout<<"选择要打开的文件：new1.txt new2.txt new3.txt"<<endl;
	 cin>>name1;
    file.open(name1,ios::in);  
	/*string contents;*/
	
  if(!file)        //检查文件是否打开成功
  {
     cout<<"Fail to open data file!"<<endl;
     return 1;
  }
  else
  {
     string temp;
     vector<string> item;   //项集的临时vector 
     int begin,end;
     while(getline(file,temp))     //一行一行读入数据
     {
        trancount++;
        begin=0; 
		 temp.erase(0,temp.find_first_not_of("\r\t\n "));   //去除字符串首部的空格
       temp.erase(temp.find_last_not_of("\r\t\n")+1); 
        while((end=temp.find('\t',begin))!=string::npos)    //每一个事务中的项是以'\t'为分隔符的
        {
           item.push_back(temp.substr(begin,end-begin));   //将每一个项插入item中
           begin=end+1;
        }
        item.push_back(temp.substr(begin));     //一个事务中的最后一项
        datavec.push_back(item);       //将一个事务中的所有项当成一个整体插入另一个大的vector中
        item.clear();   //清空item
     }
     cout<<"Press Enter to continue the processing";  //pause
     getchar();
     map<string,int> item_map; 
     for(vector<vector<string> >::size_type ix=0;ix!=datavec.size();++ix)
     {
        for(vector<string>::size_type iy=0;iy!=datavec[ix].size();++iy)
        {
         items_count[datavec[ix].at(iy)]++;    //该项集的计数加1
         item_map[datavec[ix].at(iy)]=1;       //表示该项目在该事务中存在，值为1，否则默认为0
        }
        bitmap.push_back(item_map);
        item_map.clear();      //这里一定要清空一下
     }
     map<string,int>::const_iterator map_it=items_count.begin();
     cout<<"候选1项集:"<<endl;
     while(map_it!=items_count.end())      //输出候选1项集
     {
        cout<<"{"<<map_it->first<<"}"<<endl;
        map_it++;
     }
     cout<<"Press Enter to continue the processing";  //pause
     getchar();
     map_it=items_count.begin();
     cout<<"频繁1项集(minsup=2):"<<endl;
     while(map_it!=items_count.end())          //频繁1项集
     {
        if(map_it->second>minsup)    //支持度大于2
        {
            cout.setf(ios::fixed);
			 cout<<"{"<<map_it->first<<"}"<<" 支持度:"<<setprecision(6)<<map_it->second<<endl;
            item.push_back(map_it->first);
            frequentvec.push_back(item);   //插入候选1项集的vector中
            item.clear();    
        }
        map_it++;
     }
     if(!frequentvec.empty())   //判断频繁1项集是否为空，为空则退出
     {
         cout<<"Press Enter to continue the processing";  //pause
         getchar();
         int round=1;     //生成候选项集轮次
         int found;    //是否包含有非频繁的子集,为1表示含有，有的话进行剪枝，如假设I1I4为非频繁项集，则I1I2I4要剪枝掉
         string tempstr;
         vector<string> tempvec;
         do
         {
            //生成下一轮的候选项集
            vector<vector<string> >::size_type st=frequentvec.size();
            candidatevec.clear();         //清除上一轮的候选项集
            for(vector<vector<string> >::size_type st1=0;st1<st;st1++)
            {
                for(vector<vector<string> >::size_type st2=st1+1;st2<st;st2++)
                {
                    found=0;
                    item=mergeItem(frequentvec[st1],frequentvec[st2],round);    //调用函数合并生成下一轮的候选项集
                    if(!item.empty()&&!isExist(item,candidatevec))   //若经过判断处理后返回的vector不为空且还不存在该项集，则作为候选项集加入候选vector中
                    {
                        ////////实现剪枝//////////////////////////
                        string teststr;
                        int testint;
                        tempvec=item;
                        sort(tempvec.begin(),tempvec.end());
                        while(next_permutation(tempvec.begin(),tempvec.end()))   //遍历所有的组合I1I2I4,要变成I1I4I2或其他如I2I1I4才能判断它包含I1I4这个非频繁项集
                        {
                           for(vector<string>::size_type tempst=0;tempst!=tempvec.size();tempst++) //拼接出该字符串组合
                           {
                                tempstr+=tempvec[tempst];                                
                           }
                           for(map<string,int>::const_iterator tempit=items_count.begin();tempit!=items_count.end();tempit++)
                           {
                                if(tempit->second<minsup)  //非频繁项集
                                {
                                    if(tempstr.find(tempit->first)!=string::npos)   //表示包含有非频繁子项集
                                   {
                                       found=1;
                                       teststr=tempit->first;
                                       testint=tempit->second;
                                       break;
                                   }
                                }
                           }
                           tempstr.erase();
                           if(found)   //包含非频繁子项集
                           {
                              break;
                           }
                           
                        }
                        if(!found)     //只有不包含有非频繁子项集才加入候选项集中，否则剪枝掉
                           candidatevec.push_back(item);
                       /* else
                        {
                           cout<<"剪去项集:";
                           for(vector<string>::size_type st2=0;st2!=item.size();st2++)
                             //  cout<<item[st2];
                           //cout<<" 含有非频繁子项集:"<<teststr<<" "<<testint<<"/"<<trancount<<"="<<((float)(testint)/(float)trancount);
                           //cout<<endl;
                        }*/
                        found=0;   //重置
                    }
                    
                }
            }
            frequentvec.clear();         //清除上一轮的频繁项集
            round++;
            cout<<"候选"<<round<<"项集:"<<endl;
            for(vector<vector<string> >::size_type ix=0;ix!=candidatevec.size();++ix)      //输出候选项集
            {
               cout<<"{";
               for(vector<string>::size_type iy=0;iy!=candidatevec[ix].size();++iy)
               {
                cout<<candidatevec[ix].at(iy)<<' ';
               }
               cout<<"}"<<endl;
            }
            if(candidatevec.empty())  //候选项集为空
            {
               cout<<"候选"<<round<<"项集为空!"<<endl;
            }
            int flag;    //标记某个项集在某条事务中是否出现，出现为1，不出现为0，如：{I1I2}
            int count;   //统计某个想集在整个交易的事务集中出现的次数
            string tempstr;   //临时string，用于串接各个项成一个字符串： 如: I1 I2  I3  串接为"I1I2I3"
            int mark;    //为避免执行多余的字符串串接工作
            for(vector<vector<string> >::size_type sx=0;sx!=candidatevec.size();++sx)      //构造下一轮的频繁项集
            {
                mark=1;
                count=0;
                for(vector<map<string,int> >::size_type sy=0;sy!=bitmap.size();++sy)
                {
                    flag=1;       //初始化为1，表出现
                    for(vector<string>::size_type sz=0;sz!=candidatevec[sx].size();++sz)
                    {
                       if(bitmap[sy][candidatevec[sx].at(sz)]==0)   //存在某一个子项不存在，则没出现项集
                       {
                           flag=0;
                       }
                       if(mark==1)   //只串接一次，如I1I2    否则为10个I1I2的串接
                       {
                           tempstr+=candidatevec[sx].at(sz);  //串接字符串
                       }
                    }
                    
                    if(flag)  //flag仍然为1，表示该项集在该条事务中出现了，计数加1
                    {
                       count++;
                    }
                    mark++;
                }
                
                if(count>minsup)   //支持度大于2
                {
                    frequentvec.push_back(candidatevec[sx]);        //插入频繁项集
                }
                items_count[tempstr]=count;       //对应该项集的计数值
                /////////假设此时生成的tempstr为I1I2I3,为便于后面的求置信度的计算，这里需要产生I2I1I3,I1I3I2等组合，并
                //在items_count中给它们赋予和I1I2I3相同的值
                sort(candidatevec[sx].begin(),candidatevec[sx].end());    //排序
                string tempstr2;
                while(next_permutation(candidatevec[sx].begin(),candidatevec[sx].end()))  //取下一排列组合
                {
                     for(vector<string>::size_type tempst=0;tempst!=candidatevec[sx].size();tempst++) //拼接出该字符串组合
                     {
                         tempstr2+=candidatevec[sx][tempst];
                     }
                     items_count[tempstr2]=count;  //对应该项集的计数值
                     tempstr2.erase();
                }               
                tempstr.erase();
            }
            cout<<"Press Enter to continue the processing";  //pause
            getchar();
            if(!frequentvec.empty())     //频繁项集不为空
            {
                 
                 cout<<"频繁"<<round<<"项集(minsup=2):"<<endl;
                 for(int sx=0;sx!=frequentvec.size();++sx)      //输出频繁项集
                 {
                     cout.setf(ios::fixed);
                     cout<<"{"; 
                     for(vector<string>::size_type sz=0;sz!=frequentvec[sx].size();++sz)
                     {
                        cout<<frequentvec[sx].at(sz)<<' ';
                        tempstr+=frequentvec[sx].at(sz);  //串接字符串
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
                 cout<<"没有"<<round<<"-频繁项集,Apriori算法结束!"<<endl;
             }
         }while(!frequentvec.empty());   //频繁项集不为空,则循环继续
         
        
         file.close();
      
         return 0;
     }
     else
     {         
         return 0;
     }    //end of if(!frequentvec.empty())
     
  }//end of if(!file)
}

