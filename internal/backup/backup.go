package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"os/exec"
	"github.com/Sagar1329/task-tracker-api/internal/config"
)

func RunIfNeeded(cfg *config.Config) error {
	// Ensure backup directory exists
	if err := os.MkdirAll(cfg.Backup.Directory, os.ModePerm); err != nil {
		return err
	}

	latestBackup, err := latestBackup(cfg.Backup.Directory)
	if err != nil {
		return err
	}

	// No backups found
	if latestBackup == nil {
		return createBackup(cfg)
	}

	// Backup interval has not elapsed
	if time.Since(latestBackup.ModTime()) < cfg.Backup.Interval {
		return nil
	}

	// Create a new backup
	return createBackup(cfg)
}

func latestBackup(directory string) (os.FileInfo, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var latest os.FileInfo

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		if filepath.Ext(info.Name()) != ".sql" {
			continue
		}

		if latest == nil || info.ModTime().After(latest.ModTime()) {
			latest = info
		}
	}

	if latest == nil {
		return nil, nil
	}

	return latest, nil
}

func createBackup(cfg *config.Config) error {
	fileName := fmt.Sprintf(
		"backup_%s.sql",
		time.Now().Format("20060102_150405"),
	)

	filePath := filepath.Join(cfg.Backup.Directory, fileName)

	cmd := exec.Command(
		"pg_dump",
		"-h", cfg.Database.Host,
		"-p", cfg.Database.Port,
		"-U", cfg.Database.User,
		"-d", cfg.Database.Name,
		"-f", filePath,
	)

	cmd.Env = append(
		os.Environ(),
		"PGPASSWORD="+cfg.Database.Password,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("backup failed: %v\n%s", err, string(output))
	}

	return nil
}