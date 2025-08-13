package storage

import (
	"auto-unterhaltsrechner/internal/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Storage struct {
	dataDir string
}

func New() *Storage {
	homeDir, _ := os.UserHomeDir()
	dataDir := filepath.Join(homeDir, ".auto-unterhaltsrechner")

	// Create data directory if it doesn't exist
	os.MkdirAll(dataDir, 0755)

	return &Storage{
		dataDir: dataDir,
	}
}

func (s *Storage) SaveProfile(profile *models.CarProfile) error {
	if profile == nil {
		return fmt.Errorf("profile cannot be nil")
	}

	profilesDir := filepath.Join(s.dataDir, "profiles")
	err := os.MkdirAll(profilesDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create profiles directory: %w", err)
	}

	filename := fmt.Sprintf("%s.json", profile.ID)
	filepath := filepath.Join(profilesDir, filename)

	data, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal profile: %w", err)
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write profile file: %w", err)
	}

	return nil
}

func (s *Storage) LoadProfile(id string) (*models.CarProfile, error) {
	if id == "" {
		return nil, fmt.Errorf("profile ID cannot be empty")
	}

	filename := fmt.Sprintf("%s.json", id)
	filepath := filepath.Join(s.dataDir, "profiles", filename)

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read profile file: %w", err)
	}

	var profile models.CarProfile
	err = json.Unmarshal(data, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal profile: %w", err)
	}

	return &profile, nil
}

func (s *Storage) DeleteProfile(id string) error {
	if id == "" {
		return fmt.Errorf("profile ID cannot be empty")
	}

	filename := fmt.Sprintf("%s.json", id)
	filepath := filepath.Join(s.dataDir, "profiles", filename)

	err := os.Remove(filepath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete profile file: %w", err)
	}

	return nil
}

func (s *Storage) ListProfiles() ([]*models.CarProfile, error) {
	profilesDir := filepath.Join(s.dataDir, "profiles")

	files, err := os.ReadDir(profilesDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []*models.CarProfile{}, nil
		}
		return nil, fmt.Errorf("failed to read profiles directory: %w", err)
	}

	var profiles []*models.CarProfile
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			id := file.Name()[:len(file.Name())-5] // Remove .json extension
			profile, err := s.LoadProfile(id)
			if err != nil {
				continue // Skip invalid profiles
			}
			profiles = append(profiles, profile)
		}
	}

	return profiles, nil
}

func (s *Storage) SaveSettings(settings *models.AppSettings) error {
	if settings == nil {
		return fmt.Errorf("settings cannot be nil")
	}

	filepath := filepath.Join(s.dataDir, "settings.json")

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write settings file: %w", err)
	}

	return nil
}

func (s *Storage) LoadSettings() (*models.AppSettings, error) {
	filepath := filepath.Join(s.dataDir, "settings.json")

	data, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return default settings if file doesn't exist
			return &models.AppSettings{
				Theme:                   "light",
				DefaultFuelPrice:        1.65,
				DefaultElectricityPrice: 0.35,
			}, nil
		}
		return nil, fmt.Errorf("failed to read settings file: %w", err)
	}

	var settings models.AppSettings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal settings: %w", err)
	}

	return &settings, nil
}

func (s *Storage) ExportProfileToJSON(profile *models.CarProfile, filepath string) error {
	if profile == nil {
		return fmt.Errorf("profile cannot be nil")
	}

	data, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal profile: %w", err)
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write export file: %w", err)
	}

	return nil
}

func (s *Storage) ImportProfileFromJSON(filepath string) (*models.CarProfile, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read import file: %w", err)
	}

	var profile models.CarProfile
	err = json.Unmarshal(data, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal profile: %w", err)
	}

	// Generate new ID for imported profile
	profile.ID = fmt.Sprintf("imported_%s", profile.ID)

	return &profile, nil
}
