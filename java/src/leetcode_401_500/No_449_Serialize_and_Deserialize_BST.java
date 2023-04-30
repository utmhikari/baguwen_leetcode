package leetcode_401_500;

/**
 * Created by 36191 on 2019/6/2
 */

import utils.TreeNode;

import java.util.Stack;

public class No_449_Serialize_and_Deserialize_BST {
    private void serialize(TreeNode root, StringBuilder sb) {
        if (root == null) {
            sb.append('#');
        } else {
            sb.append(root.val);
            sb.append('>');
            serialize(root.left, sb);
            sb.append('<');
            serialize(root.right, sb);
        }
    }

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();
        serialize(root, sb);
        return sb.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        char[] chars = data.toCharArray();
        int len = chars.length, status = 0;
        Stack <TreeNode> stack = new Stack<>();
        for (int i = 0; i < len; ++i) {
            char c = chars[i];
            System.out.println(c + ": " + status);
            switch (c) {
                case '>': {
                    status = 1;
                    break;
                }
                case '<': {
                    stack.pop();
                    status = 2;
                    break;
                }
                case '#': {
                    stack.push(null);
                    break;
                }
                default: {
                    int num = c - '0';
                    switch (status) {
                        case 0: {
                            stack.push(new TreeNode(num));
                            break;
                        }
                        case 1: {
                            stack.peek().left = new TreeNode(num);
                            stack.push(stack.peek().left);
                            break;
                        }
                        case 2: {
                            stack.peek().right = new TreeNode(num);
                            stack.push(stack.peek().right);
                            break;
                        }
                        default: {
                            stack.peek().val = stack.peek().val * 10 + num;
                            break;
                        }
                    }
                    status = 3;
                    break;
                }
            }
        }
        if (stack.size() > 0) {
            while (stack.size() > 1) {
                stack.pop();
            }
            return stack.pop();
        } else {
            return null;
        }
    }


    public static void main(String[] args) {
        No_449_Serialize_and_Deserialize_BST solution = new No_449_Serialize_and_Deserialize_BST();
        TreeNode root = solution.deserialize("1>#<2>#>#");
        TreeNode.output(root);
    }
}
