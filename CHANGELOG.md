# Changelog

## v0.2.0 - Major UI and UX Improvements

### New Features
- ✅ **Multilingual Support**: German/English language switching in settings
- ✅ **Improved Window Management**: Automatically centers window and sizes to 75% of screen
- ✅ **Better Performance**: Fixed splitter dragging lag issues
- ✅ **Proper Translations**: All fuel and electricity types now properly translated

### Improvements
- Better UI responsiveness with optimized container layouts
- Enhanced settings dialog with language selection
- Fixed German translations (e.g., "Öffentliche Ladestation" instead of "Public charging station")
- Cleaner code structure with proper internationalization system

### Technical
- Added comprehensive i18n system for all UI elements
- Improved data model with language-neutral enum values
- Enhanced window sizing and positioning
- Optimized container performance

### Bug Fixes
- Fixed console window appearing in background
- Resolved translation inconsistencies
- Improved splitter component performance

---

## v0.1.0 - Initial Release

### Features
- ✅ Complete vehicle cost calculation system
- ✅ Support for fuel and electric vehicles
- ✅ Multiple cost categories (fuel, electricity, tax, insurance, financing, depreciation)
- ✅ Profile management (save/load vehicle profiles)
- ✅ Export functionality (CSV, JSON, PDF)
- ✅ Vehicle comparison mode (up to 4 vehicles)
- ✅ German localization with proper number formatting
- ✅ Light/Dark theme support
- ✅ Comprehensive input validation
- ✅ Help tooltips for all input fields

### Export Features
- CSV export with detailed cost breakdown
- JSON export for profile backup/sharing
- PDF export with professional formatting (EUR currency)
- Automatic file extensions in save dialogs

### User Interface
- Clean, intuitive design with card-based layout
- Real-time calculation updates
- Responsive split-pane layout
- Proper window management (no console window)
- Professional toolbar with icons

### Technical
- Built with Go 1.21 and Fyne v2.6.2
- Windows-native executable
- No external dependencies required
- Optimized build size with stripped binaries