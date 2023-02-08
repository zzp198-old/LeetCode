package problems

//type AuthenticationToken struct {
//	tokenID     string
//	currentTime int
//	timeToLive  int
//}

type AuthenticationManager struct {
	tokens map[string]int
	ttl    int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{map[string]int{}, timeToLive}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, ok := this.tokens[tokenId]; ok && v+this.ttl > currentTime {
		this.tokens[tokenId] = currentTime
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) (ans int) {
	for _, t := range this.tokens {
		if t+this.ttl > currentTime {
			ans++
		}
	}
	return
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
