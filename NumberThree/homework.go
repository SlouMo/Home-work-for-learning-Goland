package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Кастомные ошибки
var (
	// TODO: Добавьте необходимые кастомные ошибки
	itemNonExists = errors.New("предмет отсутствует")
)

type Item interface {
	Use() (string, error)
	GetName() string
	GetWeight() float64
}

type Storable interface {
	Serialize(w io.Writer) error
	Deserialize(r io.Reader) error
}

func wrapErrorSerialize(err error) error {
	return fmt.Errorf("ошибка при сериализации: %w\n", err)
}

func wrapErrorRead(err error) error {
	return fmt.Errorf("ошибка при чтении: %w", err)
}

func wrapErrorWrite(err error) error {
	return fmt.Errorf("ошибка записи: %w", err)
}

func wrapErrorStrConvert(err error) error {
	return fmt.Errorf("ошибка преобразования строки: %w", err)
}

type Weapon struct {
	Name       string
	Damage     int
	Durability int
}

func (w *Weapon) Use() (string, error) {
	// TODO: Реализуйте возврат ошибки
	if w.Durability <= 0 {
		return fmt.Sprintf("Не удалось использовать оружие %v", w.Name), fmt.Errorf("оружие %v сломано", w.Name)
	}

	w.Durability--

	return fmt.Sprintf("Атаковали %s (%d урона)", w.Name, w.Damage), nil
}

func (w *Weapon) GetName() string {
	return w.Name
}

func (w *Weapon) GetWeight() float64 {
	return 2.5
}

func (w *Weapon) Serialize(wr io.Writer) error {
	// TODO: Реализуйте возврат ошибки
	_, err := fmt.Fprintf(wr, "Weapon|%s|%d|%d", w.Name, w.Damage, w.Durability)
	if err != nil {
		return wrapErrorSerialize(err)
	}
	return nil
}

func (w *Weapon) Deserialize(r io.Reader) error {
	// TODO: Реализуйте возврат ошибок
	data, err := io.ReadAll(r)
	if err != nil {
		return wrapErrorRead(err)
	}
	parts := strings.Split(string(data), "|")

	w.Name = parts[1]
	w.Damage, err = strconv.Atoi(parts[2])
	if err != nil {
		return wrapErrorStrConvert(err)
	}
	w.Durability, err = strconv.Atoi(parts[3])
	if err != nil {
		return wrapErrorStrConvert(err)
	}
	return nil
}

type Armor struct {
	Name    string
	Defense int
	Weight  float64
}

func (a *Armor) Use() (string, error) {
	return fmt.Sprintf("Надели %s (+%d защиты)", a.Name, a.Defense), nil
}

func (a *Armor) GetName() string {
	return a.Name
}

func (a *Armor) GetWeight() float64 {
	return a.Weight
}

func (a *Armor) Serialize(wr io.Writer) error {
	// TODO: Реализуйте возврат ошибки
	_, err := fmt.Fprintf(wr, "Armor|%s|%d|%f", a.Name, a.Defense, a.Weight)
	if err != nil {
		return wrapErrorSerialize(err)
	}
	return nil
}

func (a *Armor) Deserialize(r io.Reader) error {
	// TODO: Реализуйте возврат ошибок
	data, err := io.ReadAll(r)
	if err != nil {
		return wrapErrorRead(err)
	}
	parts := strings.Split(string(data), "|")

	a.Name = parts[1]
	a.Defense, err = strconv.Atoi(parts[2])
	if err != nil {
		return wrapErrorStrConvert(err)
	}
	a.Weight, err = strconv.ParseFloat(parts[3], 64)
	if err != nil {
		return wrapErrorStrConvert(err)
	}
	return nil
}

type Potion struct {
	Name    string
	Effect  string
	Charges int
}

func (p *Potion) Use() (string, error) {
	// TODO: Реализуйте возврат ошибки
	if p.Charges <= 0 {
		return fmt.Sprintf("Не удалось использовать зелье %v", p.Name), fmt.Errorf("зелье %v закончилось", p.Name)
	}

	p.Charges--

	return fmt.Sprintf("Использовали %s (%s)", p.Name, p.Effect), nil
}

func (p *Potion) GetName() string {
	return p.Name
}

func (p *Potion) GetWeight() float64 {
	return 0.5
}

func DescribeItem(i Item) (string, error) {
	// TODO: Реализуйте возврат ошибки
	if i == nil {
		return "", itemNonExists
	}

	return fmt.Sprintf("%s (вес: %.1f)", i.GetName(), i.GetWeight()), nil
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T

	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T any, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))

	for i, item := range items {
		result[i] = transform(item)
	}

	return result
}

func Find[T any](items []T, condition func(T) bool) (T, bool) {
	for _, item := range items {
		if condition(item) {
			return item, true
		}
	}

	var zero T

	return zero, false
}

type Inventory struct {
	Items []Item
}

func (inv *Inventory) AddItem(item Item) error {
	// TODO: Проверка на nil
	if item == nil {
		return fmt.Errorf("пойди туда — не знаю куда, добавь то — не знаю что.\n Ты добавляешь nil в инвентарь")
	}
	inv.Items = append(inv.Items, item)
	return nil
}

func (inv *Inventory) GetWeapons() []*Weapon {
	weapons := Filter(inv.Items, func(item Item) bool {
		_, ok := item.(*Weapon)
		return ok
	})

	return Map(weapons, func(item Item) *Weapon {
		return item.(*Weapon)
	})
}

func (inv *Inventory) GetBrokenItems() []Item {
	return Filter(inv.Items, func(item Item) bool {
		switch v := item.(type) {
		case *Weapon:
			return v.Durability <= 0
		case *Potion:
			return v.Charges <= 0
		default:
			return false
		}
	})
}

func (inv *Inventory) GetItemNames() []string {
	return Map(inv.Items, func(item Item) string {
		return item.GetName()
	})
}

func (inv *Inventory) FindItemByName(name string) (Item, bool) {
	return Find(inv.Items, func(item Item) bool {
		return item.GetName() == name
	})
}

func (inv *Inventory) Save(w io.Writer) error {
	// TODO: Реализуйте возврат ошибки
	for _, item := range inv.Items {
		if storable, ok := item.(Storable); ok {
			err := storable.Serialize(w)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(w)
			if err != nil {
				return wrapErrorWrite(err)
			}
		}
	}
	return nil
}

func (inv *Inventory) Load(r io.Reader) error {
	// TODO: Реализуйте возврат ошибки
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Weapon") {
			var w Weapon

			r := strings.NewReader(line)

			err := w.Deserialize(r)
			if err != nil {
				return err
			}

			err = inv.AddItem(&w)
			if err != nil {
				return err
			}
		} else if strings.HasPrefix(line, "Armor") {
			var a Armor

			r := strings.NewReader(line)

			err := a.Deserialize(r)
			if err != nil {
				return err
			}

			err = inv.AddItem(&a)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SafeUse(item Item) (result string, err error) {
	// TODO: Используйте defer с recover для перехвата паники
	// TODO: Для оружия с именем "Ящик Пандоры" вызовите панику
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Был открыт Ящик Пандоры, но силы Света смогли его закрыть:", r)
		}
	}()

	if item == nil {
		return "", fmt.Errorf("нельзя использовать то, чего нет")
	}

	if item.GetName() == "Ящик Пандоры" {
		panic("Ты открыл Ящик Пандоры!\nБеги глупец, но тебе уже не спастись!")
	}
	return item.Use()
}

func main() {
	// TODO: Реализуйте логику/вызовы:
	// TODO: 1. Обработку ошибок везде
	// TODO: 2. Use предмета до потери прочности и обработку ошибки при потере прочности
	// TODO: 3. DescribeItem с предметом и с nil
	// TODO: 4. Обработку ошибок сохранения/загрузки в файл
	// TODO: 5. Обработку паники для "Ящика Пандоры"
	inv := Inventory{}

	sword := &Weapon{Name: "Меч", Damage: 10, Durability: 5}
	healthPotion := &Potion{Name: "Лечебное", Effect: "+50 HP", Charges: 3}
	pandoraBox := &Weapon{Name: "Ящик Пандоры", Damage: math.MaxInt, Durability: math.MaxInt}

	items := []Item{sword, healthPotion, pandoraBox, nil}
	fmt.Println("Добавление в инвентарь")
	for _, item := range items {
		err := inv.AddItem(item)
		if err != nil {
			fmt.Printf("Проблема при добавлении в инвентарь: %v\n", err)
		}
	}
	fmt.Println("--------------")
	fmt.Println("Вывод описания")
	for _, item := range items {
		describe, err := DescribeItem(item)
		if err != nil {
			fmt.Printf("Проблема при получении описания: %v\n", err)
		} else {
			fmt.Println(describe)
		}
	}
	fmt.Println("--------------")
	fmt.Println("Безопасное использование")
	for _, item := range items {
		action, err := SafeUse(item)
		if err != nil {
			fmt.Printf("Проблема при использовании: %v\n", err)
		} else if action != "" {
			fmt.Println(action)
		}
	}
	fmt.Println("--------------")
	fmt.Println("Сохраняем в файл")

	file, err := os.OpenFile("homework_solved.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
	}

	err = inv.Save(file)
	if err != nil {
		fmt.Printf("Ошибка при сохранении: %v\n", err)
	}

	fmt.Println("Ломаем файл")

	_, err = fmt.Fprintf(file, "Weapon||")
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
	}

	fmt.Println("Загружаем из файла")
	inv = Inventory{}

	file, err = os.Open("homework_solved.txt")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
	}

	err = inv.Load(file)
	if err != nil {
		fmt.Printf("Ошибка загрузки файла: %v\n", err)
	}

	names := inv.GetItemNames()

	fmt.Println("\nИмена предметов:", names)

	for _, item := range inv.Items {
		describe, err := DescribeItem(item)
		if err != nil {
			fmt.Printf("Проблема при получении описания: %v\n", err)
		} else {
			fmt.Println("-", describe)
		}
	}
}
