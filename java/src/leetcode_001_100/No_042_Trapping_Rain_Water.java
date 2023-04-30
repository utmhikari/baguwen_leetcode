package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/13
 * 042. 蓄雨池
 */

public class No_042_Trapping_Rain_Water {

    /*
    public int min(int a, int b) {
        return a < b ? a : b;
    }

    public ArrayList<Integer> getBarriers(int[] height, ArrayList<Integer> heightIndices) {
        ArrayList<Integer> originalBarrierIndices = new ArrayList<>();
        int minHeight = 0;
        int index = 0;
        while(index < heightIndices.size()) {
            if(height[heightIndices.get(index)] >= minHeight) {
                originalBarrierIndices.add(heightIndices.get(index));
                int s = originalBarrierIndices.size();
                if(s != 1) {
                    minHeight = min(height[heightIndices.get(index)], height[originalBarrierIndices.get(s - 2)]);
                }
            }
            minHeight =
            index++;
        }

        System.out.print("OriBarrierIndices: ");
        for(int i : originalBarrierIndices) {
            System.out.print(i + ", ");
        }
        System.out.println();

        ArrayList<Integer> reducedIndices = new ArrayList<>();
        ArrayList<Integer> startContiIndices = new ArrayList<>(); // conti indices for barrierIndices
        for(int i = 0; i < originalBarrierIndices.size(); i++) {
            if(i == 0) {
                startContiIndices.add(i);
            } else {
                if(originalBarrierIndices.get(i) != originalBarrierIndices.get(i - 1) + 1) {
                    startContiIndices.add(i);
                }
            }
        }

        System.out.print("ContiIndices: ");
        for(int i : startContiIndices) {
            System.out.print(i + ", ");
        }
        System.out.println();

        for(int i = 0; i <= startContiIndices.size() - 1; i++) {
            int maxHeight = 0;
            int maxIndex = -1;
            for(int j = startContiIndices.get(i);
                j < (i == startContiIndices.size() - 1 ? originalBarrierIndices.size() : startContiIndices.get(i + 1)); j++) {
                if(height[originalBarrierIndices.get(j)] >= maxHeight) {
                    maxHeight = height[originalBarrierIndices.get(j)];
                    maxIndex = originalBarrierIndices.get(j);
                }
            }
            reducedIndices.add(maxIndex);
        }
        return reducedIndices;
    }

    public int trap(int[] height) {
        int len = height.length;
        if(len <= 2) {
            return 0;
        }
        ArrayList<Integer> barrierIndices = new ArrayList<>();
        for(int i = 0; i < len; i++) {
            barrierIndices.add(i);
        }
        barrierIndices = getBarriers(height, barrierIndices);

        System.out.print("BarrierIndices: ");
        for(int i = 0; i < barrierIndices.size(); i++) {
            System.out.print(barrierIndices.get(i) + ", ");
        }
        System.out.println();

        int sumSize = 0;
        for(int i = 0; i < barrierIndices.size() - 1; i++) {
            int startIndex = barrierIndices.get(i);
            int endIndex = barrierIndices.get(i + 1);
            int h = min(height[startIndex], height[endIndex]);
            sumSize += h * (endIndex - startIndex - 1);
            for(int j = startIndex + 1; j < endIndex; j++) {
                sumSize -= height[j];
            }
        }
        return sumSize;
    }
    */

    public int trap(int[] A) {
        if (A.length < 3) return 0;

        int ans = 0;
        int l = 0, r = A.length - 1;

        // find the left and right edge which can hold water
        while (l < r && A[l] <= A[l + 1]) l++;
        while (l < r && A[r] <= A[r - 1]) r--;

        while (l < r) {
            int left = A[l];
            int right = A[r];
            if (left <= right) {
                // add volume until an edge larger than the left edge
                while (l < r && left >= A[++l]) {
                    ans += left - A[l];
                }
            } else {
                // add volume until an edge larger than the right volume
                while (l < r && A[--r] <= right) {
                    ans += right - A[r];
                }
            }
        }
        return ans;
    }

    public static void main(String[] args) {
        No_042_Trapping_Rain_Water solution = new No_042_Trapping_Rain_Water();
        System.out.println(solution.trap(new int[] {0,1,0,2,1,0,1,3,2,1,2,1}));
        System.out.println(solution.trap(new int[] {4, 2, 3}));
        System.out.println(solution.trap(new int[] {2, 1, 2, 4, 4, 4, 4, 2, 1, 2}));
    }
}
