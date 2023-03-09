package problems

func minimumRecolors(blocks string, k int) int {
    W,B:=0,0
    for i:=0;i<k;i++{
        if blocks[i]=='W'{
            W++
        }
        if blocks[i]=='B'{
            B++
        }
    }
    ans:=W
    for i:=k;i<len(blocks);i++{
        if blocks[i]=='W'{
            W++
        }
    if blocks[i]=='B'{
            B++
        }
          if blocks[i-k]=='W'{
            W--
        }
          if blocks[i-k]=='B'{
            B--
        }
        if W<ans{
            ans=W
        }
    }
return ans
}
