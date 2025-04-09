package common

import "testing"

func Test_Queue(t *testing.T){
	t.Run("works with FiFo", func(t *testing.T) {
		q := Queue[string]{}
		q.Enqueue("Thomas")
		q.Enqueue("Maia")
		q.Enqueue("Amyzona")
		q.Enqueue("Jadeeee")

		got, _ := q.Dequeue()
		want := "Thomas"
		
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		for i := 0; i < 2; i++ {
			q.Dequeue()
		}

		got, _ = q.Dequeue()
		want = "Jadeeee"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		_, err := q.Dequeue()

		if err == nil {
			t.Errorf("got nil, want %v", err)
		}
	})

	t.Run("FiFo logic should work for Queue of only one element", func(t *testing.T) {
		q := Queue[string]{}
		q.Enqueue("Thomas")
		
		_, err := q.Dequeue()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}

		_, err  = q.Dequeue()
		if err == nil {
			t.Errorf("got nil, want %v", err)
		}
	})
}

