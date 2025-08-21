Go
package main

import (
	"sync"
)

// AutomationScript represents a single automation script
type AutomationScript struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Script    string `json:"script"`
	Trigger   string `json:"trigger"`
	Interval  int    `json:"interval"`
	Enabled   bool   `json:"enabled"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// AutomationScriptIntegrator represents the automation script integrator
type AutomationScriptIntegrator struct {
	scripts     map[string]AutomationScript
	scriptMutex sync.RWMutex
}

// NewAutomationScriptIntegrator returns a new instance of AutomationScriptIntegrator
func NewAutomationScriptIntegrator() *AutomationScriptIntegrator {
	return &AutomationScriptIntegrator{
		scripts: make(map[string]AutomationScript),
	}
}

// AddScript adds a new automation script to the integrator
func (asi *AutomationScriptIntegrator) AddScript(script AutomationScript) {
	asi.scriptMutex.Lock()
	defer asi.scriptMutex.Unlock()
	asi.scripts[script.ID] = script
}

// GetScript returns an automation script by ID
func (asi *AutomationScriptIntegrator) GetScript(id string) (AutomationScript, bool) {
	asi.scriptMutex.RLock()
	defer asi.scriptMutex.RUnlock()
	script, ok := asi.scripts[id]
	return script, ok
}

// ListScripts returns a list of all automation scripts
func (asi *AutomationScriptIntegrator) ListScripts() []AutomationScript {
	asi.scriptMutex.RLock()
	defer asi.scriptMutex.RUnlock()
	scripts := make([]AutomationScript, 0, len(asi.scripts))
	for _, script := range asi.scripts {
		scripts = append(scripts, script)
	}
	return scripts
}

// RemoveScript removes an automation script by ID
func (asi *AutomationScriptIntegrator) RemoveScript(id string) {
	asi.scriptMutex.Lock()
	defer asi.scriptMutex.Unlock()
	delete(asi.scripts, id)
}