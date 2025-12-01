package main

import (
	"fmt"
	"sync"
)

func inputStream(n int, out chan<- int) {
	defer close(out)
	for i := 1; i <= n; i++ {
		out <- i
	}
}
func doubler(in <-chan int, out chan<- int) {
	defer close(out)
	for v := range in {
		out <- v * 2
	}
}

func printer(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("Резултат%d\n", v)
	}
}

func main() {
	in := make(chan int)
	mid := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go inputStream(10, in)
	go doubler(in, mid)
	go printer(mid, &wg)

	wg.Wait()
}

/*//sync нужен для того что бы main ждал пока printer дочитает канал и завершил свою работу -
	// и что б он не завершился раньше времени
	// in первый этап передачи данных (
	in - для чисел 1..N
	mid - для удвоения чисел
	in := make (chan int)
	mid := make (chan int) - Соиденение горутин

	2.1 Запус генератора чисел
	go inputStream(10, in)
	Горутина начинает генерировать чсила  от 1 до 10 и отправляет в канал in
	2.2 Запуск удвоения
	go doubler(in, mid)
	Это горутина:
	1. Читает числа in
	2. Умножает чиса на 2
	3. отправляет рузультат в канал mid
	 Когда in закроется и числа закончатся она сама завершится и закроет mid

	3. Printer работает в основном потоке
	printer(mid)

	Printer читает числа из mid , печатает результат Х и выходит когда мид закрывается
goo run main.go
//)*/

/*
func producer(ch chan int){
	for i := 1; i <= 5; i++{
		fmt.Println("отправляю")
		ch <- i
	}
	close(ch)
	fmt.Println("Producer закончил работу ")
}

func consumer(ch chan int){
	for v := range ch {
		fmt.Println("Получил", v)
	}
	fmt.Println("Consumer закончил работать ")

}

func main(){
	ch := make(chan int)
	producer(ch)
	producer(ch)

	fmt.Println("main завершен ")
} */
