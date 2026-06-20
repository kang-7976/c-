import subprocess
import os

def check_phone_info():
    print("=" * 40)
    print("📱 手机硬件快速体检报告")
    print("=" * 40)

    # 定义我们要查询的属性列表
    props = {
        "手机型号": "ro.product.model",
        "品牌": "ro.product.brand",
        "安卓版本": "ro.build.version.release",
        "SDK版本": "ro.build.version.sdk",
        "CPU架构": "ro.product.cpu.abi",
        "屏幕密度(DPI)": "ro.sf.lcd_density",
        "语言区域": "ro.product.locale"
    }

    for name, key in props.items():
        try:
            # 使用 getprop 命令获取属性值
            result = subprocess.run(
                ["getprop", key],
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                text=True
            )
            value = result.stdout.strip()
            if not value:
                value = "无法获取"
            print(f"{name:<12}: {value}")
        except Exception as e:
            print(f"{name:<12}: 读取失败 ({e})")

    print("=" * 40)
    print("✅ 检测完成！")

if __name__ == "__main__":
    check_phone_info()

