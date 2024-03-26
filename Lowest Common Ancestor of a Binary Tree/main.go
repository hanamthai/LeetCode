/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /**
 Ta bắt đầu tìm kiếm từ nút gốc của cây và xem xét các trường hợp cơ bản:

Nếu nút gốc là nil hoặc nút gốc trùng với nút p hoặc q, ta trả về nút gốc đó.
Trong trường hợp này, nút gốc sẽ là tổ tiên chung gần nhất của p và q.
Sau đó, ta tiếp tục tìm kiếm ở cả hai phía trái và phải của nút gốc:

Nếu p và q nằm ở các phía khác nhau của nút gốc, tức là một nút nằm bên trái và một nút nằm bên phải của nút gốc, thì nút gốc chính là tổ tiên chung gần nhất của p và q, do đó ta trả về nút gốc đó.
Nếu cả hai left và right đều không nil, nghĩa là p và q đều nằm ở các phía khác nhau của nút gốc, ta trả về nút gốc đó.
Nếu chỉ một trong left hoặc right không nil, ta trả về nút đó, vì nó chứa một trong hai nút p hoặc q.

Cuối cùng, ta trả về kết quả của nút đầu tiên mà đồng thời chứa p và q, hoặc là tổ tiên của p và q.
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}