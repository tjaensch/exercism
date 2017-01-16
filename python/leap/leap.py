def leap_year(year):
    if year%100 == 0 and year%400 == 0:
    	return true
    elif year%100 == 0:
        return false
    elif year%4 == 0:
        return true
    else:
        return false