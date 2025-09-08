from pypinyin import lazy_pinyin, Style

def generate_keyword_line(text):
    pinyin_list = lazy_pinyin(text, style=Style.TONE3, neutral_tone_with_five=True)
    processed_pinyin = [py.rstrip('12345') for py in pinyin_list]
    pinyin_str = ' '.join(processed_pinyin)
    return f'{pinyin_str} @{text}'

# 生成新唤醒词
wake_words = ['小助手', '开始工作', '星期五']
for word in wake_words:
    print(generate_keyword_line(word))
