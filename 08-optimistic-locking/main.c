#include <stdio.h>
#include <stdatomic.h>
#include <stdbool.h>
#include <pthread.h>

atomic_int count = ATOMIC_VAR_INIT(0);

void* incrementCount(void* arg) {
    int oldValue = atomic_load(&count);
    int newValue = oldValue + 1;
    if (!atomic_compare_exchange_strong(&count, &oldValue, newValue)) {
        printf("inc failed\n");
    }
    pthread_exit(NULL);
}

int main() {
    pthread_t thread1, thread2;
    pthread_create(&thread1, NULL, incrementCount, NULL);
    pthread_create(&thread2, NULL, incrementCount, NULL);

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    return 0;
}
