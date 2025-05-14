SERVICES := gateway auth-service subnoddit-service
PRIMARY := config-server

run:
	@echo "==> Starting primary service: $(PRIMARY)"
	@make -C $(PRIMARY) run & \
	PRIMARY_PID=$$!; \
	sleep 10; \
	echo "==> Starting other services..."; \
	for svc in $(SERVICES); do \
		echo "==> Starting $$svc"; \
		make -C $$svc run & \
	done; \
	trap "echo 'Shutting down...'; kill $$PRIMARY_PID; kill 0" SIGINT SIGTERM; \
	wait
