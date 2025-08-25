package main

import (
	"fmt"
	"io"
	"strings"
)

// TODO: Реализуйте структуры:
// TODO: - Weapon: Name (string), Damage (int), Durability (int)
// TODO: - Armor: Name (string), Defense (int), Weight (float64)
// TODO: - Potion: Name (string), Effect (string), Charges (int)
// TODO:
// TODO: Можете добавить свои структуры :)
// TODO:
// TODO: Для каждой структуры реализуйте:
// TODO: - Метод Use() string (описание использования, например "Используется <имя>", и изменение Durability или Charges и т.д.)
// TODO: - Методы интерфейса Item
type Weapon struct {
	Name       string
	Damage     int
	Durability int
}

func (weapon *Weapon) Use() string {
	if weapon.Durability <= 0 {
		return "Ты не можешь использовать сломанное оружие"
	}

	weapon.Durability--
	return fmt.Sprintf("Ты использовал %v и нанес %v урона", weapon.Name, weapon.Damage)
}

func (weapon *Weapon) GetName() string {
	return weapon.Name
}

func (weapon *Weapon) GetWeight() float64 {
	return float64(weapon.Durability)*0.25 + 1
}

type Armor struct {
	Name       string
	Defense    int
	Durability int
	Weight     float64
}

func (armor *Armor) Use() string {
	if armor.Durability <= 0 {
		return "Твоя броня сломана! Ты гол, как сокол!"
	}

	armor.Durability--
	return fmt.Sprintf("Урон снижен на %v единиц", armor.Defense)
}

func (armor *Armor) GetName() string {
	return armor.Name
}

func (armor *Armor) GetWeight() float64 {
	return armor.Weight
}

type Potion struct {
	Name    string
	Effect  string
	Charges int
}

func (potion *Potion) Use() string {
	if potion.Charges <= 0 {
		return "Отсутствуют заряды"
	}

	potion.Charges--
	return fmt.Sprintf("Эффект \"%v\" наложен", potion.Effect)
}

func (potion *Potion) GetName() string {
	return potion.Name
}

func (potion *Potion) GetWeight() float64 {
	return float64(potion.Charges) * 0.5
}

type Item interface {
	GetName() string
	GetWeight() float64
	Use() string
}

// TODO: Реализуйте функцию
func DescribeItem(i Item) string {
	// Функция должна возвращать:
	// - "Предмет отсутствует" если i == nil
	// - "<название> (вес: <вес>)" в остальных случаях
	if i == nil {
		return "Предмет отсутствует"
	}
	return fmt.Sprintf("%v (вес: %.2f)", i.GetName(), i.GetWeight())
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	// TODO: Верните новый слайс только с элементами, для которых predicate вернул true
	var result []T

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, R any](items []T, transform func(T) R) []R {
	// TODO: Примените transform к каждому элементу и верните слайс с результатами
	var result []R
	for _, item := range items {
		result = append(result, transform(item))
	}
	return result
}

func Find[T any](items []T, condition func(T) bool) (T, bool) {
	// TODO: Найдите первый элемент, удовлетворяющий condition и верните элемент и true или zero value и false
	var zero T
	for _, item := range items {
		if condition(item) {
			return item, true
		}
	}
	return zero, false
}

type Inventory struct {
	Items []Item
}

func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
}

func (inv *Inventory) GetWeapons() []*Weapon {
	// TODO: Используйте Filter для отбора Weapon, затем Map для преобразования Item -> *Weapon
	filterValue := Filter(inv.Items, func(item Item) bool {
		_, ok := item.(*Weapon)
		return ok
	})

	return Map(filterValue, func(item Item) *Weapon {
		value, _ := item.(*Weapon)
		return value
	})
}

func (inv *Inventory) GetBrokenItems() []Item {
	// TODO: Используйте Filter для отбора:
	// TODO: - Weapon: Durability <= 0
	// TODO: - Potion: Charges <= 0
	// TODO:
	// TODO: Подсказка: поможет приведение типов - item.(type)
	return Filter(inv.Items, func(item Item) bool {
		if value, ok := item.(*Weapon); ok {
			if value.Durability <= 0 {
				return true
			}
			return false
		}
		if value, ok := item.(*Armor); ok {
			if value.Durability <= 0 {
				return true
			}
			return false
		}
		if value, ok := item.(*Potion); ok {
			if value.Charges <= 0 {
				return true
			}
			return false
		}
		return false
	})
}

func (inv *Inventory) GetItemNames() []string {
	// TODO: Используйте Map для преобразования []Item -> []string
	return Map(inv.Items, func(item Item) string {
		return item.GetName()
	})
}

func (inv *Inventory) FindItemByName(name string) (Item, bool) {
	// TODO: Используйте Find для поиска по имени
	return Find(inv.Items, func(item Item) bool {
		if strings.Contains(item.GetName(), name) {
			return true
		}
		return false
	})
}

// TODO: Бонус: реализуйте интефейс Storable для Weapon и Armor:
// TODO: - Weapon: формат "Weapon|Name|Damage|Durability"
// TODO: - Armor: формат "Armor|Name|Defense|Weight"

type Storable interface {
	Serialize(w io.Writer)
	Deserialize(r io.Reader)
}

func (inv *Inventory) Save(w io.Writer) {
	// TODO: Бонус: сделайте сохранение/загрузку инвентаря в/из файла
}

func (inv *Inventory) Load(r io.Reader) {
	// TODO: Бонус: сделайте сохранение/загрузку инвентаря в/из файла
}

func main() {
	// TODO: Создайте инвентарь и добавьте:
	// TODO: - Оружие: "Меч" (урон 10, прочность 5)
	// TODO: - Броню: "Щит" (защита 5, вес 4.5)
	// TODO: - Зелье: "Лечебное" (+50 HP, 3 заряда)
	// TODO: - Оружие: "Сломанный лук" (урон 5, прочность 0)
	inventory := Inventory{}
	sword := &Weapon{
		Name:       "Меч короля Артура",
		Damage:     10,
		Durability: 5,
	}
	shield := &Armor{
		Name:       "Щит Ахилла",
		Defense:    5,
		Durability: 5,
		Weight:     4.5,
	}
	potion := &Potion{
		Name:    "Лечебное зелье",
		Effect:  "+50 HP",
		Charges: 3,
	}
	brokenBow := &Weapon{
		Name:       "Сломанный лук",
		Damage:     5,
		Durability: 0,
	}
	inventory.AddItem(sword)
	inventory.AddItem(shield)
	inventory.AddItem(potion)
	inventory.AddItem(brokenBow)

	// TODO: Реализуйте логику/вызовы:
	// TODO: 1. Use предмета с выводом в консоль
	fmt.Println("1. Use предмета с выводом в консоль")
	for _, item := range inventory.Items {
		fmt.Println(item.Use())
	}
	fmt.Println()
	// TODO: 2. DescribeItem с предметом и с nil, так же с выводом в консоль
	fmt.Println("2. DescribeItem с предметом и с nil, так же с выводом в консоль")
	for index := range inventory.Items {
		fmt.Println(DescribeItem(inventory.Items[index]))
	}
	fmt.Println(DescribeItem(nil))
	fmt.Println()
	// TODO: 3. Вывести в консоль результат вызова GetWeapons (должны вернуться только меч и лук)
	fmt.Println("3. Вывести в консоль результат вызова GetWeapons (должны вернуться только меч и лук)")
	weapons := inventory.GetWeapons()
	weaponsStr := make([]string, len(weapons))
	for _, item := range weapons {
		weaponsStr = append(weaponsStr, item.GetName())
	}
	fmt.Printf("%v\n", weaponsStr)
	fmt.Println()
	// TODO: 4. Вывести в консоль результат вызова GetBrokenItems (должен вернуть сломанный лук)
	fmt.Println("4. Вывести в консоль результат вызова GetBrokenItems (должен вернуть сломанный лук)")
	weaponsBroken := inventory.GetBrokenItems()
	weaponsBrokenStr := make([]string, len(weaponsBroken))
	for _, item := range weaponsBroken {
		weaponsBrokenStr = append(weaponsBrokenStr, item.GetName())
	}
	fmt.Printf("%v\n", weaponsBrokenStr)
	fmt.Println()
	// TODO: 5. Вывести в консоль результат вызова GetItemNames (все названия)
	fmt.Println("5. Вывести в консоль результат вызова GetItemNames (все названия)")
	fmt.Println(inventory.GetItemNames())
	fmt.Println()
	// TODO: 6. Вывести в консоль результат вызова FindItemByName (поиск "Щит")
	fmt.Println("6. Вывести в консоль результат вызова FindItemByName (поиск \"Щит\")")
	fmt.Println(inventory.FindItemByName("Щит"))
	fmt.Println()
	// TODO: Бонус: сделайте сохранение инвентаря в файл и загрузку инвентаря из файла
}
