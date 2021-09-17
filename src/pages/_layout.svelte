<script>
  import {
    Header,
    HeaderUtilities,
    HeaderAction,
    HeaderGlobalAction,
    HeaderPanelLinks,
    HeaderPanelDivider,
    HeaderPanelLink,
    HeaderNav,
    HeaderNavItem,
    HeaderNavMenu,
    SideNav,
    SideNavItems,
    SideNavMenu,
    SideNavMenuItem,
    SideNavLink,
    SideNavDivider,
    SkipToContent,
    Content,
    Grid,
    Row,
    Column,
  } from "carbon-components-svelte"
  import Fade16 from "carbon-icons-svelte/lib/Fade16"
  import SettingsAdjust20 from "carbon-icons-svelte/lib/SettingsAdjust20"
  import CanDecorator from "../decorators/Can.svelte"

  let isSideNavOpen = true
  let isOpen = false
</script>

<Header
  company="Analytic Service"
  platformName="Ril engine platform"
  bind:isSideNavOpen
>
  <div slot="skip-to-content">
    <SkipToContent />
  </div>

  <HeaderNav>
    <HeaderNavItem href="/" text="Link 1" />
    <HeaderNavItem href="/" text="Link 2" />
    <HeaderNavItem href="/" text="Link 3" />
    <HeaderNavMenu text="Menu">
      <HeaderNavItem href="/" text="Link 1" />
      <HeaderNavItem href="/" text="Link 2" />
      <HeaderNavItem href="/" text="Link 3" />
    </HeaderNavMenu>
    <HeaderNavItem href="/" text="Link 4" />
  </HeaderNav>

  <HeaderUtilities>
    <HeaderGlobalAction aria-label="Settings" icon={SettingsAdjust20} />
    <HeaderAction bind:isOpen>
      <HeaderPanelLinks>
        <HeaderPanelDivider>Switcher subject 1</HeaderPanelDivider>
        <HeaderPanelLink>Switcher item 1</HeaderPanelLink>
        <HeaderPanelDivider>Switcher subject 2</HeaderPanelDivider>
        <HeaderPanelLink>Switcher item 1</HeaderPanelLink>
        <HeaderPanelLink>Switcher item 2</HeaderPanelLink>
        <HeaderPanelLink>Switcher item 3</HeaderPanelLink>
        <HeaderPanelLink>Switcher item 4</HeaderPanelLink>
        <HeaderPanelLink>Switcher item 5</HeaderPanelLink>
      </HeaderPanelLinks>
    </HeaderAction>
  </HeaderUtilities>

  <SideNav bind:isOpen={isSideNavOpen} rail>
    <SideNavItems>
      <SideNavLink icon={Fade16} text="users" href="/users" isSelected />
      <SideNavLink icon={Fade16} text="dummies" href="/dummies" />
      <SideNavLink icon={Fade16} text="Link 3" href="/" />
      <SideNavMenu icon={Fade16} text="权限管理">
        <SideNavMenuItem href="/" text="菜单管理" />
        <SideNavMenuItem href="/" text="Link 2" />
        <SideNavMenuItem href="/"
          ><div class="flex items-center">
            <Fade16 /> Some Text
          </div></SideNavMenuItem
        >
      </SideNavMenu>
      <SideNavDivider />
      <SideNavLink icon={Fade16} text="Link 4" href="/" />
    </SideNavItems>
  </SideNav>
</Header>

<Content>
  <slot decorator={CanDecorator} />
</Content>
