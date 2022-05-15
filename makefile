check:
	echo "checkin"
build_wizard:
	cd frontend && npm run build_wizard
	# cp frontend/public/build/plug_simplewizard.js backend/src/stdplugs/simplewizard/exec_script.js && \
	# cp frontend/public/build/plug_simplewizard.css backend/src/stdplugs/simplewizard/exec_style.css
	# @echo "Needs server restart"
build_dashed:
	cd frontend && npm run build_dashed