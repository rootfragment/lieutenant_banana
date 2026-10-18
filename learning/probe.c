//go:build ignore
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

char LICENSE[] SEC("license") = "Dual BSD/GPL";

struct event {
    u32 pid;
    u32 uid;
    u8 comm[16];
};

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 256 * 1024);
} events SEC(".maps");

// Makes bpf2go emit a Go type for struct event
const struct event *unused __attribute__((unused));

SEC("tracepoint/sched/sched_process_exec")
int handle_exec(void *ctx)
{
    char comm[16];
    bpf_get_current_comm(&comm, sizeof(comm));

    // Only care about "ls"
    if (comm[0] != 'l' || comm[1] != 's' || comm[2] != '\0')
        return 0;

    struct event *e = bpf_ringbuf_reserve(&events, sizeof(*e), 0);
    if (!e)
        return 0;

    e->pid = bpf_get_current_pid_tgid() >> 32;
    e->uid = (u32)bpf_get_current_uid_gid();
    __builtin_memcpy(e->comm, comm, sizeof(comm));

    bpf_ringbuf_submit(e, 0);
    return 0;
}
