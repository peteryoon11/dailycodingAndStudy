
# #reference : http://navercast.naver.com/contents.nhn?rid=2871&contents_id=85293 

import copy #목적지와 도착지를 설정해준다 

departure = 'home'
destination = 'school' 
print ("-----------[", departure, "->", destination,"]----------" )
#① 지도상의 모든 건물들과 집에서 각 건물들까지의 최단 거리를 나타내는 표를 만든다. 
landscape = { 'home': {'hairShop':5, 'superMarket':10, 'EnglishAcademy':9}, 'hairShop' : {'home':5 ,'superMarket':3, 'bank':11}, 'superMarket' : {'hairShop':3, 'home':10, 'EnglishAcademy':7, 'restourant':3}, 'EnglishAcademy': {'home':9, 'superMarket':7, 'school':12}, 'restourant' : {'superMarket':3, 'bank':4}, 'bank' : {'hairShop':11, 'restourant':4, 'EnglishAcademy':7, 'school':2}, 'school' : {'bank':2, 'EnglishAcademy':12} } 
routing = {} 
for place in landscape.keys(): 
    routing[place]={'shortestDist':0, 'route':[], 'visited':0} 
#④ 
def visitPlace(visit): 
    routing[visit]['visited'] = 1 
    for toGo, betweenDist in landscape[visit].items(): 
        toDist = routing[visit]['shortestDist'] + betweenDist 
        if (routing[toGo]['shortestDist'] >= toDist) or not routing[toGo]['route']: 
            routing[toGo]['shortestDist'] = toDist 
            routing[toGo]['route'] = copy.deepcopy(routing[visit]['route']) 
            routing[toGo]['route'].append(visit) 
            #② 집과 직접 길로 이어진 건물들까지의 최단 거리는 지도에 표시된 값으로 적고 그렇지 않은 건물들은 빈 칸으로 놓아둔다. 여기서 빈 칸의 값은 무한대를 뜻한다. 
            visitPlace(departure)
            # ③ 거리가 가장 짧은 건물부터 긴 건물 순으로 방문하고 방문한 건물은 색깔로 칠해 구별한다. 이때 방문한 경로도 색칠한다. ④ 새로운 건물을 방문하면 그 건물과 이어진 건물들까지의 거리를 새로 바꾼다. 단, 이전에 이미 최단 거리가 구해졌었다면 거리를 서로 비교해 작은 것으로 바꾸거나 유지한다. ⑤ 그래프의 모든 건물들을 방문할 때까지 ③,④의 과정을 반복한다. "
            # "" 
            while 1 : #③ 
                minDist = max(routing.values(), key=lambda x:x['shortestDist'])['shortestDist'] 
                toVisit = '' 
                for name, search in routing.items(): 
                    if 0 < search['shortestDist'] <= minDist and not search['visited']: 
                        minDist = search['shortestDist'] 
                        toVisit = name 
                        #⑤ 
                        if toVisit == '': 
                            break 
                        #④ 
                        visitPlace(toVisit) 
                        print ("["+toVisit+"]" )
                        print ("Dist :", minDist )
                        print ("\n", "[", departure, "->", destination,"]" )
                        print ("Route : ", routing[destination]['route'] )
                        print ("ShortestDistance : ", routing[destination]['shortestDist'] )

# 출처: http://phaphaya.tistory.com/34 [pAPaYA]

