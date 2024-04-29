import os

# Путь к директории
directory = "C:\\Users\\alexe\\Desktop\\programing\\bmstu_go\\loops\\"

# # Получаем список файлов и директорий в указанной директории
# file_list = os.listdir(directory)

# # Фильтруем только файлы
# file_list = [f for f in file_list if os.path.isfile(os.path.join(directory, f))]

# # Выводим список файлов
# for file_name in file_list:
#     print(file_name)

with open(directory + 'loops.go') as f:
    a = f.read()
    old = new = 'a'
    while old and new:
        with open(directory + 'loops.go', 'w') as w:
            old = input()
            new = input()
            if not (old or new):
                w.write(a)
                break
            if new in a and (new not in old or a.count(new) > a.count(old)):
                print('hueta')
                w.write(a)
            else:
                w.write(a.replace(old, new))