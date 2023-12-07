# Create a product and price for 3 items

menu_items = [
    "Espresso", "Cappuccino", "Latte", "Mocha", "Americano",
    "Macchiato", "Flat White", "Cortado", "Affogato", "Irish Coffee",
    "Cold Brew", "Iced Latte", "Chai Latte", "Matcha Latte", "Turmeric Latte",
    "Hot Chocolate", "White Chocolate Mocha", "Caramel Macchiato", "Vanilla Latte", "Hazelnut Espresso",
    "Croissant", "Bagel with Cream Cheese", "Avocado Toast", "Greek Yogurt Parfait", "Fruit Salad",
    "Blueberry Muffin", "Chocolate Chip Cookie", "Almond Croissant", "Cinnamon Roll", "Cheese Danish",
    "Spinach and Feta Quiche", "Ham and Cheese Croissant", "Smashed Avo Baguette", "Turkey Club Sandwich", "Caprese Panini",
    "Vegetarian Wrap", "Chicken Caesar Salad", "Smoked Salmon Bagel", "Egg Salad Sandwich", "BLT Sandwich",
    "Classic Pancakes", "French Toast", "Belgian Waffles", "Eggs Benedict", "Vegan Pancake Stack",
    "Frittata", "Oatmeal with Berries", "Granola Parfait", "Acai Bowl", "Quinoa Salad",
    "Caprese Skewers", "Stuffed Mushrooms", "Bruschetta", "Garlic Bread", "Truffle Fries",
    "Margherita Pizza", "Pepperoni Pizza", "Vegetarian Pizza", "BBQ Chicken Pizza", "Pesto Pasta",
    "Spaghetti Bolognese", "Chicken Alfredo", "Shrimp Scampi", "Vegetable Stir-Fry", "Teriyaki Salmon",
    "Grilled Chicken Salad", "Caesar Salad", "Greek Salad", "Cobb Salad", "Mango Tango Smoothie",
    "Strawberry Banana Smoothie", "Green Detox Juice", "Mango Passionfruit Iced Tea", "Raspberry Lemonade", "Sparkling Water",
    "Chocolate Brownie Sundae", "New York Cheesecake", "Tiramisu", "Apple Pie", "Chocolate Lava Cake",
    "Fruit Sorbet", "Pistachio Gelato", "Affogato", "Raspberry Chocolate Truffle", "Creme Brulee",
    "Artisanal Cheese Platter", "Charcuterie Board", "Mixed Nuts and Olives", "Brussels Sprouts with Balsamic Glaze", "Stuffed Bell Peppers"
]

menu_prices = {
    "Espresso": 2.50, "Cappuccino": 3.50, "Latte": 4.00, "Mocha": 4.50, "Americano": 3.00,
    "Macchiato": 3.50, "Flat White": 4.00, "Cortado": 3.50, "Affogato": 5.50, "Irish Coffee": 6.00,
    "Cold Brew": 3.50, "Iced Latte": 4.50, "Chai Latte": 4.50, "Matcha Latte": 5.00, "Turmeric Latte": 5.00,
    "Hot Chocolate": 4.00, "White Chocolate Mocha": 4.50, "Caramel Macchiato": 4.50, "Vanilla Latte": 4.50, "Hazelnut Espresso": 4.00,
    "Croissant": 2.50, "Bagel with Cream Cheese": 3.00, "Avocado Toast": 5.00, "Greek Yogurt Parfait": 4.50, "Fruit Salad": 4.00,
    "Blueberry Muffin": 2.50, "Chocolate Chip Cookie": 2.00, "Almond Croissant": 3.00, "Cinnamon Roll": 4.00, "Cheese Danish": 3.50,
    "Spinach and Feta Quiche": 5.00, "Ham and Cheese Croissant": 3.50, "Smashed Avo Baguette": 6.00, "Turkey Club Sandwich": 7.00, "Caprese Panini": 6.50,
    "Vegetarian Wrap": 6.00, "Chicken Caesar Salad": 8.00, "Smoked Salmon Bagel": 7.50, "Egg Salad Sandwich": 6.50, "BLT Sandwich": 7.00,
    "Classic Pancakes": 5.50, "French Toast": 6.00, "Belgian Waffles": 5.50, "Eggs Benedict": 8.00, "Vegan Pancake Stack": 6.50,
    "Frittata": 7.00, "Oatmeal with Berries": 4.50, "Granola Parfait": 5.00, "Acai Bowl": 7.50, "Quinoa Salad": 6.50,
    "Caprese Skewers": 5.00, "Stuffed Mushrooms": 6.00, "Bruschetta": 4.50, "Garlic Bread": 3.50, "Truffle Fries": 6.00,
    "Margherita Pizza": 8.00, "Pepperoni Pizza": 9.00, "Vegetarian Pizza": 8.50, "BBQ Chicken Pizza": 9.50, "Pesto Pasta": 7.50,
    "Spaghetti Bolognese": 8.00, "Chicken Alfredo": 9.00, "Shrimp Scampi": 10.00, "Vegetable Stir-Fry": 7.50, "Teriyaki Salmon": 11.00,
    "Grilled Chicken Salad": 8.50, "Caesar Salad": 7.50, "Greek Salad": 8.00, "Cobb Salad": 9.00, "Mango Tango Smoothie": 5.00,
    "Strawberry Banana Smoothie": 4.50, "Green Detox Juice": 6.00, "Mango Passionfruit Iced Tea": 3.50, "Raspberry Lemonade": 4.00, "Sparkling Water": 2.50,
    "Chocolate Brownie Sundae": 6.50, "New York Cheesecake": 5.00, "Tiramisu": 7.00, "Apple Pie": 5.50, "Chocolate Lava Cake": 6.00,
    "Fruit Sorbet": 4.50, "Pistachio Gelato": 5.00, "Affogato": 5.50, "Raspberry Chocolate Truffle": 3.50, "Creme Brulee": 7.50,
    "Artisanal Cheese Platter": 10.00, "Charcuterie Board": 12.00, "Mixed Nuts and Olives": 6.50, "Brussels Sprouts with Balsamic Glaze": 5.00, "Stuffed Bell Peppers": 8.00
}

for i, item in enumerate(menu_items, start=1):
    price = menu_prices[item]
    exec(f"p{i}_name = '{item}'")
    exec(f"p{i}_price = {price}")
    # print(f"{eval('p' + str(i) + '_name')}: €{eval('p' + str(i) + '_price'):.2f}")

# Company
company_name = "Bean Bliss Café"
company_address = '''123 Mocha Street
                Java Town
                Roastville
                Coffee County
                Caffeine
                12345
                (555) 987-6543
                feedback@beanblisscafe.com'''

# Footnote
message = '''Thank you for choosing Bean Bliss Café!
                We hope you enjoyed your coffee experience with us.
                Your satisfaction is our priority.
                If you have any feedback or suggestions,
                please feel free to reach out.
                
                Follow us on social media:
                @BeanBlissCafe #BeanBlissCafe'''

# Create a top border
print("*" * 50)
print("\t\t{}".format(company_name.title()))
print("\t\t{}".format(company_address))

print()
print()
print("\t\t{} \t\t{}".format("Product", "Price (€)"))
print()

total = 0.0
tax_rate = 0.23
cash = 650.00

for i, item in enumerate(menu_items, start=1):
    price = menu_prices[item]
    total = total + price
    print("\t\t{} \t\t{:.2f}".format(item, price))

print()

print("\t\t{} \t\t{:.2f}".format("SUBTOTAL", total))
print("\t\t{} \t\t{:.2f}".format("TAX @ 23%", total * tax_rate))
print("\t\t{} \t\t{:.2f}".format("TOTAL TO PAY", total + (total * tax_rate)))

print()

print("\t\t{}".format("PAID BY"))
print("\t\t\t{} \t\t{:.2f}".format("CASH", cash))
print("\t\t\t{} \t\t{:.2f}".format("CHANGE", cash - (total + (total * tax_rate))))

print()

# Create end border:
print("\t\t{}".format(message))
print("*" * 50)
