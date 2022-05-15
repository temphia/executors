stdplug_compile_simplewizard:
	cd frontend && npm run build_stdplug_simplewizard
	# cp frontend/public/build/plug_simplewizard.js backend/src/stdplugs/simplewizard/exec_script.js && \
	# cp frontend/public/build/plug_simplewizard.css backend/src/stdplugs/simplewizard/exec_style.css
	# @echo "Needs server restart"
stdplug_compile_simpledash:
	cd frontend && npm run build_stdplug_simpledash