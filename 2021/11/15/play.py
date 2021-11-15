import random


restaurantArray = [
    "배달의민족-진천족보-족발", 
    "배달의민족-도담그루-중국집", 
    "스시현-초밥", 
    "슬로우먼데이-버거", 
    "이마트-푸드코트", 
    ]



recommedNumber = 2 

todayChoiceNumbers = random.sample(range(0, len(restaurantArray)),recommedNumber)

print("추천 가게")
for i in todayChoiceNumbers:
    print("result "+restaurantArray[i])



