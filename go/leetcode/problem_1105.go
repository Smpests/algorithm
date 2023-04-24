package leetcode

// 题目链接：https://leetcode.cn/problems/filling-bookcase-shelves/
func minHeightShelves(books [][]int, shelfWidth int) int {
    bookCount := len(books)
    // 增加一个初始状态位，即空书架
    states := make([]int, bookCount + 1)
    // 初始化状态为每本书都单独放一层书架后的高度
    for i:= 1; i <= bookCount; i++ {
        states[i] = states[i - 1] + books[i - 1][1]
    }

    for i := 0; i < bookCount; i++ {  // 放入第i本书
        floorMaxHeight, leftWidth := 0, shelfWidth
        for j := i; j > -1; j-- {  // 枚举第i本书与i前面的书可能的放置情况
            leftWidth -= books[j][0]
            if leftWidth < 0 {
                break  // 因为顺序放，因此枚举到超出宽度时退出
            }
            if books[j][1] > floorMaxHeight {
                floorMaxHeight = books[j][1]  // 层高为层内最高的书
            }
            newStates := states[j] + floorMaxHeight
            if newStates < states[i + 1] {
                states[i + 1] = newStates  // 更新状态，因为0状态无书，因此+1
            } 
        }
    }
    return states[bookCount]
}

// func minHeightShelves(books [][]int, shelfWidth int) int {
//     bookCount := len(books)
//     states := make([]int, bookCount)  // 存放每本书放下后的最小高度
//     states[0] = books[0][1]  // 第一本书一定会放在第一层
//     for i := 1; i < bookCount; i++ {
//         // 初始化为每层单独放一本书时的高度
//         states[i] = states[i - 1] + books[i][1]
//     }

//     for i := 1; i < bookCount; i++ {
//         floorMaxHeight, leftWidth := 0, shelfWidth
//         for j := i; j > -1; j-- {
//             leftWidth -= books[j][0]
//             if leftWidth < 0 {
//                 break
//             }
//             if books[j][1] > floorMaxHeight {
//                 floorMaxHeight = books[j][1]
//             }
//             niceTry := floorMaxHeight
//             if j > 0 {
//                 niceTry += states[j - 1]
//             }
//             if states[i] > niceTry {
//                 states[i] = niceTry
//             }
//         }
//     }
//     return states[bookCount - 1]
// }

// 超时，配合LRUCache结构存储dfs的输入输出可通过
// func minHeightShelves(books [][]int, shelfWidth int) int {
//     var dfs func(int) int
//     dfs = func(index int) int {
//         if index < 0 {
//             return 0
//         }
//         result, floorMaxHeight, leftWidth := 1000000, 0, shelfWidth
//         for i := index; i > -1; i-- {
//             leftWidth -= books[i][0]
//             if leftWidth < 0 {
//                 break
//             }
//             if books[i][1] > floorMaxHeight {
//                 floorMaxHeight = books[i][1]
//             }
//             next := dfs(i - 1) + floorMaxHeight
//             if next < result {
//                 result = next
//             }
//         }
//         return result
//     }
//     return dfs(len(books) - 1)
// }