export function await1000() {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve("test")
        }, 1000)
    })
}

export function await5000() {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve("test")
        }, 5000)
    })
}
