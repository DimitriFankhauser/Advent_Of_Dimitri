def calculate(num1: int, num2: int, operator: str) -> float:
    if operator == '+':
        return float(num1 + num2)
    if operator == '%':
        return float(num2%num1)
    if operator == '-':
        return float(num1 - num2)
    if operator == '*':
        return float(num1 * num2)
    if operator == '/':
        return float(num1 / num2)
    if operator == '>':
        if num1 > num2:
            return float(num2)
    if operator == '<':
        if num1 < num2:
            return float(num1)
