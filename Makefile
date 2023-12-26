
define make_win_tmp
	mkdir -p tmp
	wget -O tmp/opengl32.7z https://downloads.fdossena.com/geth.php?r=mesa64-latest 
	cd tmp && 7z e opengl32.7z
endef

define windows_env
	GOOS=windows \
	GGO_ENABLED=1 \
	CC=/usr/bin/x86_64-w64-mingw32-gcc \
	CXX=/usr/bin/x86_64-w64-mingw32-c++ \
	$1	
endef

define compile_windows
	$(call windows_env, fyne package --os windows --release --src ./cmd/EduTrack/)
	mkdir -p builds
	mv ./cmd/EduTrack/EduTrack.exe ./builds/EduTrack.exe
endef

define compile_linux
	fyne package --os linux --release --src ./cmd/EduTrack/
	mkdir -p builds
	mv EduTrack.tar.xz builds/EduTrack-linux64.tar.xz
endef


clean:
	rm -rf builds tmp windows-zip
	rm -rf ./cmd/EduTrack/EduTrack
	rm -rf ./cmd/EduTrack/EduTrack.exe
	rm -rf ./cmd/Installer/EduTrack\ Installer.exe
	rm -rf ./cmd/Installer/EduTrack\ Installer

user-install:
	$(call compile_linux)
	cd builds && \
	tar -xvf EduTrack-linux64.tar.xz && \
	make user-install

install:
	$(call compile_linux)
	cd builds && \
	tar -xvf EduTrack-linux64.tar.xz && \
	sudo make install


uninstall:
	$(call compile_linux)
	cd builds && \
	tar -xvf EduTrack-linux64.tar.xz && \
	sudo make uninstall
	rm -rf ~/.config/EduTrack/

user-uninstall:
	$(call compile_linux)
	cd builds && \
	tar -xvf EduTrack-linux64.tar.xz && \
	make user-uninstall
	rm -rf ~/.config/EduTrack/

build-to-windows:
	$(call compile_windows)

build-to-linux:
	$(call compile_linux)

windows-zip:
	$(call make_win_tmp)
	mkdir -p windows-zip
	cp builds/EduTrack.exe windows-zip
	cp tmp/opengl32.dll windows-zip
	cp README.md windows-zip
	cp tmp/README.txt windows-zip/opengl-README.txt
	cp screenshots windows-zip/ -rf
	zip -r builds/EduTrack-win64.zip windows-zip
	rm -rf windows-zip

windows-installer:
	mkdir -p builds
	$(call make_win_tmp)
	$(call compile_windows)
	cp builds/EduTrack.exe ./internal/installer/install/files/EduTrack.exe -rf
	cp ./tmp/opengl32.dll ./internal/installer/install/files/opengl32.dll -rf
	$(call windows_env,fyne package --os windows --release --src ./cmd/Installer/)
	mv ./cmd/Installer/EduTrack\ Installer.exe ./builds/EduTrack-Installer-win64.exe

linux-installer:
	mkdir -p builds
	$(call compile_linux)
	cp ./builds/EduTrack-linux64.tar.xz ./internal/installer/install/files/EduTrack-linux64.tar.xz -rf
	go build -o builds/EduTrack-Installer-linux64 ./cmd/Installer/main_linux.go
